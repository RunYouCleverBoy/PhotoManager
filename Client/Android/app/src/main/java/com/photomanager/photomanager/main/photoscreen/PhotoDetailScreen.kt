package com.photomanager.photomanager.main.photoscreen

import androidx.compose.foundation.layout.Arrangement
import androidx.compose.foundation.layout.Column
import androidx.compose.foundation.layout.fillMaxSize
import androidx.compose.runtime.Composable
import androidx.compose.ui.Modifier
import androidx.compose.ui.res.stringResource
import com.photomanager.photomanager.R
import me.saket.telephoto.zoomable.coil.ZoomableAsyncImage

@Composable
fun PhotoDetailScreen(uri: String) {
    Column(modifier = Modifier.fillMaxSize(), verticalArrangement = Arrangement.Center) {
        ZoomableAsyncImage(
            modifier = Modifier.fillMaxSize(),
            model = uri,
            contentDescription = stringResource(id = R.string.zoomed_image),
        )
    }
}