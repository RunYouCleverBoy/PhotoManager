package com.photomanager.photomanager.main.components

import android.net.Uri
import android.os.Build
import android.os.ext.SdkExtensions
import android.provider.MediaStore
import androidx.activity.compose.rememberLauncherForActivityResult
import androidx.activity.result.PickVisualMediaRequest
import androidx.activity.result.contract.ActivityResultContracts
import androidx.compose.runtime.Composable
import androidx.compose.runtime.LaunchedEffect
import androidx.compose.runtime.derivedStateOf
import androidx.compose.runtime.getValue
import androidx.compose.runtime.remember

@Composable
fun ImagePicker(onPick: (List<Uri>?) -> Unit) {
    val maxImages by remember {
        derivedStateOf {
            if (SdkExtensions.getExtensionVersion(Build.VERSION_CODES.R) >= 2) {
                MediaStore.getPickImagesMaxLimit()
            } else {
                1
            }
        }
    }
    val pickMedia = rememberLauncherForActivityResult(ActivityResultContracts.PickMultipleVisualMedia(maxImages)) { picks ->
        // Handle the returned uri
        onPick(picks)
    }

    LaunchedEffect(Unit) {
        pickMedia.launch(PickVisualMediaRequest(ActivityResultContracts.PickVisualMedia.ImageOnly))
    }
}