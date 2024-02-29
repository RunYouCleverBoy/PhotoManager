package com.photomanager.photomanager.mvi

import androidx.lifecycle.ViewModel
import kotlinx.coroutines.flow.MutableSharedFlow
import kotlinx.coroutines.flow.MutableStateFlow
import kotlinx.coroutines.flow.SharedFlow
import kotlinx.coroutines.flow.StateFlow
import timber.log.Timber

abstract class MVIViewModel<STATE, EVENT, ACTION>(initialState: STATE) : ViewModel()  {
    protected val stateMutable = MutableStateFlow(initialState)
    val state = stateMutable as StateFlow<STATE>

    private val _action = MutableSharedFlow<ACTION>(extraBufferCapacity = 10)
    val action = _action as SharedFlow<ACTION>

    protected fun emit(action: ACTION) {
        val success = _action.tryEmit(action)
        if (!success) {
            Timber.w("Failed to emit action: $action")
        }
    }

    abstract fun dispatchEvent(event: EVENT)
}