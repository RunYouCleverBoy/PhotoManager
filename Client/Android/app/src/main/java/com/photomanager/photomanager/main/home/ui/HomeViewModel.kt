package com.photomanager.photomanager.main.home.ui

import android.net.Uri
import androidx.lifecycle.viewModelScope
import com.photomanager.photomanager.main.home.LazyBulk
import com.photomanager.photomanager.main.home.model.SearchCriteria
import com.photomanager.photomanager.main.home.model.WorkflowStage
import com.photomanager.photomanager.main.home.repository.PhotoRepo
import com.photomanager.photomanager.main.nav.MainNavPath
import com.photomanager.photomanager.mvi.MVIViewModel
import dagger.hilt.android.lifecycle.HiltViewModel
import kotlinx.coroutines.Dispatchers
import kotlinx.coroutines.flow.update
import kotlinx.coroutines.launch
import kotlinx.coroutines.withContext
import timber.log.Timber
import java.text.SimpleDateFormat
import java.util.Date
import java.util.GregorianCalendar
import java.util.Locale
import javax.inject.Inject

@HiltViewModel
class HomeViewModel @Inject constructor(
    private val photoRepo: PhotoRepo,
    private val homeTabRepo: HomeTabRepo
) : MVIViewModel<HomeState, HomeEvent, HomeAction>(HomeState()) {
    private val dateFormatter = SimpleDateFormat("dd/MM/yyyy", Locale.ROOT)
    private val defaultPhotoDate = Date(GregorianCalendar(2000, 1, 1).timeInMillis)

    init {
        stateMutable.update { state ->
            state.copy(
                tabDescriptors = homeTabRepo.getTabs(),
                currentMode = WorkflowStage.FOOTAGE,
                footageSearchCriteria = SearchCriteria(stage = WorkflowStage.FOOTAGE),
                collectionSearchCriteria = SearchCriteria(stage = WorkflowStage.COLLECTION)
            )
        }
    }

    override fun dispatchEvent(event: HomeEvent) {
        Timber.d("dispatchEvent: $event")
        when (event) {
            is HomeEvent.OnImageClicked -> emit(HomeAction.OpenImage(event.uri))
            is HomeEvent.OnImagesPicked -> addToFootage(event.uris)
            is HomeEvent.OnAddToCollection -> addToCollection(event.ids)
            is HomeEvent.OnTabSelected -> onTabSelected(event.stage)
            is HomeEvent.OnFootageAboutToMiss -> onMiss(event.index, WorkflowStage.FOOTAGE)
            is HomeEvent.OnCollectionAboutToMiss -> onMiss(event.index, WorkflowStage.COLLECTION)
        }
    }

    private fun onTabSelected(stage: WorkflowStage) {
        if (state.value.isBusy) return
        markBusy(true)
        viewModelScope.launch {
            refreshFromDb(stage)
            emit(
                HomeAction.NavigateTo(
                    when (stage) {
                        WorkflowStage.FOOTAGE -> MainNavPath.HomeFootage
                        WorkflowStage.COLLECTION -> MainNavPath.HomeCollection
                    }
                )
            )
        }.invokeOnCompletion { markBusy(false) }
    }

    private fun addToCollection(ids: List<String>) {
        markBusy(true)
        viewModelScope.launch {
            photoRepo.addToCollection(ids)
        }.invokeOnCompletion { markBusy(false) }
    }

    private fun addToFootage(uris: List<Uri>) {
        markBusy(true)
        viewModelScope.launch {
            photoRepo.importPhotos(uris, WorkflowStage.FOOTAGE)
            refreshFromDb(WorkflowStage.FOOTAGE)
        }.invokeOnCompletion { markBusy(false) }
    }

    private fun onMiss(index: Int, flowStage: WorkflowStage) {
        if (state.value.isBusy) return
        markBusy(true)
        val range = IntRange(index, index + PAGE_SIZE)
        val searchCriteria = when (flowStage) {
            WorkflowStage.FOOTAGE -> state.value.footageSearchCriteria
            WorkflowStage.COLLECTION -> state.value.collectionSearchCriteria
        }

        viewModelScope.launch {
            photoRepo.getPhotosByCriteria(searchCriteria, range).collect { lst ->
                val newItems = withContext(Dispatchers.IO) {
                    lst.map { photo ->
                        ImageUIDescriptor.Data(
                            photo.id,
                            Uri.parse(photo.url),
                            photo.metadata.description?.takeIf { it.isNotBlank() }
                                ?: dateFormatter.format(
                                    photo.metadata.shotDate ?: defaultPhotoDate
                                ))
                    }
                }

                populateDataWithNewItems(index, flowStage, newItems)
            }
        }.invokeOnCompletion {
            markBusy(false)
        }
    }

    private fun populateDataWithNewItems(
        startingIndex: Int,
        flowStage: WorkflowStage,
        newItems: List<ImageUIDescriptor.Data>
    ) {
        stateMutable.update { state ->
            when (flowStage) {
                WorkflowStage.FOOTAGE -> state.copy(
                    footage = state.footage.copy(
                        withAddedData = newItems,
                        startingIndex = startingIndex
                    )
                )

                WorkflowStage.COLLECTION -> state.copy(
                    collection = state.collection.copy(
                        withAddedData = newItems,
                        startingIndex = startingIndex
                    )
                )
            }
        }
    }

    private suspend fun refreshFromDb(stage: WorkflowStage) {
        val count = photoRepo.getSize(
            when (stage) {
                WorkflowStage.FOOTAGE -> state.value.footageSearchCriteria
                WorkflowStage.COLLECTION -> state.value.collectionSearchCriteria
            }
        )
        val newLazyBulk = LazyBulk<ImageUIDescriptor>(
            totalRunSize = count,
            cachedSize = 1000
        ) { ImageUIDescriptor.Loading }
        stateMutable.update { state ->
            when (stage) {
                WorkflowStage.FOOTAGE -> {
                    state.copy(
                        currentMode = stage,
                        footage = newLazyBulk
                    )
                }

                WorkflowStage.COLLECTION -> state.copy(
                    currentMode = stage,
                    collection = newLazyBulk
                )
            }
        }
    }

    private fun markBusy(busy: Boolean) {
        stateMutable.update { state -> state.copy(isBusy = busy) }
    }

    companion object {
        private const val PAGE_SIZE = 300
    }
}

