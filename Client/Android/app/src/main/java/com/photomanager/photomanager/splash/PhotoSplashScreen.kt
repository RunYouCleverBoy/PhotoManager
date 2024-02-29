package com.photomanager.photomanager.splash

import androidx.compose.animation.core.EaseOutBounce
import androidx.compose.animation.core.animateFloatAsState
import androidx.compose.animation.core.tween
import androidx.compose.foundation.Image
import androidx.compose.foundation.clickable
import androidx.compose.foundation.layout.Column
import androidx.compose.foundation.layout.fillMaxSize
import androidx.compose.foundation.layout.size
import androidx.compose.runtime.Composable
import androidx.compose.runtime.LaunchedEffect
import androidx.compose.runtime.collectAsState
import androidx.compose.runtime.getValue
import androidx.compose.runtime.mutableIntStateOf
import androidx.compose.runtime.remember
import androidx.compose.runtime.setValue
import androidx.compose.ui.Alignment
import androidx.compose.ui.Modifier
import androidx.compose.ui.graphics.graphicsLayer
import androidx.compose.ui.layout.onGloballyPositioned
import androidx.compose.ui.res.painterResource
import androidx.compose.ui.res.stringResource
import androidx.compose.ui.unit.dp
import androidx.hilt.navigation.compose.hiltViewModel
import com.photomanager.photomanager.R

@Composable
fun PhotoSplashScreen(onComplete: () -> Unit) {
    val vm = hiltViewModel<SplashViewModel>()
    val state: SplashState by vm.state.collectAsState()
    val animation by animateFloatAsState(
        targetValue = state.animationTarget,
        animationSpec = tween(durationMillis = 500, easing = EaseOutBounce),
        label = "SplashAnimation"
    ) {
        vm.dispatchEvent(SplashEvent.OnAnimationComplete)
    }

    LaunchedEffect(Unit){
        vm.dispatchEvent(SplashEvent.OnSplashScreenOpened)
        vm.action.collect{
            when(it){
                is SplashAction.NavigateToMain -> onComplete()
            }
        }
    }

    var height by remember{ mutableIntStateOf(0) }
    Column(
        modifier = Modifier
            .fillMaxSize()
            .onGloballyPositioned { coordinates ->
                height = coordinates.size.height
            }
        ,
        horizontalAlignment = Alignment.CenterHorizontally
    ) {
        Image(
            modifier = Modifier
                .size(300.dp)
                .graphicsLayer {
                    translationY = animation * (height - 300.dp.toPx())
                }
                .clickable { vm.dispatchEvent(SplashEvent.OnImageClicked) },
            painter = painterResource(id = R.drawable.hedgehog_01_13),
            contentDescription = stringResource(
                id = R.string.splash_screen_content_description
            )
        )
    }
}