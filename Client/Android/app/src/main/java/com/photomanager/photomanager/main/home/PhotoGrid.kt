package com.photomanager.photomanager.main.home

import android.net.Uri
import androidx.compose.foundation.background
import androidx.compose.foundation.clickable
import androidx.compose.foundation.layout.Box
import androidx.compose.foundation.layout.Column
import androidx.compose.foundation.layout.aspectRatio
import androidx.compose.foundation.layout.fillMaxSize
import androidx.compose.foundation.layout.fillMaxWidth
import androidx.compose.foundation.layout.height
import androidx.compose.foundation.lazy.grid.GridCells
import androidx.compose.foundation.lazy.grid.LazyVerticalGrid
import androidx.compose.material3.Text
import androidx.compose.runtime.Composable
import androidx.compose.runtime.LaunchedEffect
import androidx.compose.ui.Alignment
import androidx.compose.ui.Modifier
import androidx.compose.ui.graphics.Color
import androidx.compose.ui.res.stringResource
import androidx.compose.ui.unit.dp
import coil.compose.AsyncImage
import com.photomanager.photomanager.R
import com.photomanager.photomanager.main.home.ui.ImageUIDescriptor
import timber.log.Timber

@Composable
fun PhotoGrid(
    imagesForGrid: LazyBulk<ImageUIDescriptor>,
    pageSize : Int = 100,
    onApproachingWindowEnd: (Int) -> Unit = {},
    onImageClicked: (Uri) -> Unit = {}
) {
    LazyVerticalGrid(modifier = Modifier.fillMaxSize(), columns = GridCells.Fixed(3)) {
        items(imagesForGrid.totalRunSize) { index ->
            val imageDescriptor: ImageUIDescriptor = imagesForGrid[index]
            Timber.d("Image descriptor $imageDescriptor")
            LaunchedEffect(index / (pageSize / 4)) {
                if (imagesForGrid.lookAhead(index = index, peekSize = pageSize / 2)) {
                    return@LaunchedEffect
                }
                onApproachingWindowEnd(index)
            }

            OneCell2(imageDescriptor, onImageClicked)
        }
    }
}

@Composable
fun OneCell2(imageDescriptor: ImageUIDescriptor, onImageClicked: (Uri) -> Unit) {
    when (imageDescriptor) {
        is ImageUIDescriptor.Data -> DataCell(imageDescriptor, onImageClicked)
        ImageUIDescriptor.Loading -> PlaceHolder()
    }
}

@Composable
fun DataCell(imageDescriptor: ImageUIDescriptor.Data, onImageClicked: (Uri) -> Unit) {
    val uri = imageDescriptor.uri
    Column(modifier = Modifier
        .fillMaxWidth()
        .height(200.dp).clickable {
            onImageClicked(uri)
        }, horizontalAlignment = Alignment.CenterHorizontally) {
        AsyncImage(
            modifier = Modifier.fillMaxWidth().aspectRatio(1f),
            model = uri, contentDescription = stringResource(id = R.string.grid_content_description, 0))
        Text(text = imageDescriptor.caption, modifier = Modifier.fillMaxWidth())
    }
}

@Composable
private fun PlaceHolder() {
    Box(modifier = Modifier
        .fillMaxWidth()
        .aspectRatio(1f)
        .background(Color.Gray), contentAlignment = Alignment.Center) {
        Text(stringResource(id = R.string.loading))
    }
}