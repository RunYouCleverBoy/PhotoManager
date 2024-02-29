package com.photomanager.photomanager.main.photoscreen

import androidx.compose.foundation.layout.Box
import androidx.compose.foundation.layout.fillMaxSize
import androidx.compose.material3.Text
import androidx.compose.runtime.Composable
import androidx.compose.ui.Modifier

@Composable
fun PhotoDetailScreen(uri: String) {
    Box(modifier = Modifier.fillMaxSize()) {
        Text(text = "Photo Detail Screen for\n$uri")
    }
}