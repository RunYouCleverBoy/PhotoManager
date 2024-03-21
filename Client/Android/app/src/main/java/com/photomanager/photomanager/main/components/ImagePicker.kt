package com.photomanager.photomanager.main.components

import android.content.Intent
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
import androidx.compose.ui.platform.LocalContext

@Composable
fun ImagePicker(persistUris: Boolean = true, onPick: (List<Uri>?) -> Unit) {
    val maxImages by remember {
        derivedStateOf {
            if (SdkExtensions.getExtensionVersion(Build.VERSION_CODES.R) >= 2) {
                MediaStore.getPickImagesMaxLimit()
            } else {
                100
            }
        }
    }

    val context = LocalContext.current
    val contentResolver by remember { derivedStateOf { context.contentResolver } }
    val pickMedia = rememberLauncherForActivityResult(ActivityResultContracts.PickMultipleVisualMedia(maxImages)) { picks ->
        // Handle the returned uris
        if (persistUris) {
            picks.forEach { uri ->
                val flagGrantReadUriPermission = Intent.FLAG_GRANT_READ_URI_PERMISSION
                contentResolver.takePersistableUriPermission(uri, flagGrantReadUriPermission)
            }
        }

        onPick(picks)
    }

    LaunchedEffect(Unit) {
        pickMedia.launch(PickVisualMediaRequest(ActivityResultContracts.PickVisualMedia.ImageOnly))
    }
}