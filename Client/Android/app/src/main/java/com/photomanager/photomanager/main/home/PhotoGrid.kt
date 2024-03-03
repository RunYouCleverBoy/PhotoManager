package com.photomanager.photomanager.main.home

import android.net.Uri
import androidx.compose.foundation.clickable
import androidx.compose.foundation.layout.Column
import androidx.compose.foundation.layout.fillMaxSize
import androidx.compose.foundation.layout.fillMaxWidth
import androidx.compose.foundation.layout.height
import androidx.compose.foundation.lazy.grid.GridCells
import androidx.compose.foundation.lazy.grid.LazyVerticalGrid
import androidx.compose.material3.Text
import androidx.compose.runtime.Composable
import androidx.compose.ui.Alignment
import androidx.compose.ui.Modifier
import androidx.compose.ui.platform.LocalContext
import androidx.compose.ui.res.stringResource
import androidx.compose.ui.unit.dp
import coil.compose.AsyncImage
import coil.request.ImageRequest
import com.photomanager.photomanager.R
import timber.log.Timber

@Composable
fun PhotoGrid(
    imagesForGrid: List<ImageUIDescriptor>,
    onApproachingWindowEnd: (Int) -> Unit = {},
    onImageClicked: (Uri) -> Unit = {}
) {
    LazyVerticalGrid(modifier = Modifier.fillMaxSize(), columns = GridCells.Fixed(3)) {
        items(imagesForGrid.size) { index ->
            val (uri, caption) = imagesForGrid[index]
            Column(
                modifier = Modifier
                    .fillMaxWidth()
                    .height(200.dp),
                horizontalAlignment = Alignment.CenterHorizontally
            ) {
                AsyncImage(
                    modifier = Modifier
                        .fillMaxWidth()
                        .weight(1f)
                        .clickable {
                            Timber.d("Clicked on image $uri")
                            onImageClicked(uri)
                        },
                    model = ImageRequest.Builder(LocalContext.current).data(uri).build(),
                    contentDescription = stringResource(
                        id = R.string.grid_content_description,
                        index
                    )
                )
                Text(text = caption, modifier = Modifier.fillMaxWidth())
            }
        }
    }
}