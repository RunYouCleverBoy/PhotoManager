package com.photomanager.photomanager.splash

import com.photomanager.photomanager.mvi.MVIViewModel
import kotlinx.coroutines.flow.update

class SplashViewModel : MVIViewModel<SplashState, SplashEvent, SplashAction>(SplashState()) {
    override fun dispatchEvent(event: SplashEvent) {
        // no-op
        when (event) {
            SplashEvent.OnAnimationComplete -> {
                emit(SplashAction.NavigateToMain)
            }
            SplashEvent.OnImageClicked -> {
                stateMutable.update { it.copy(animationTarget = 1f - it.animationTarget) }
            }

            else -> stateMutable.update { it.copy(animationTarget = 1f) }
        }
    }
}

data class SplashState(val animationTarget: Float = 0f)
sealed class SplashEvent {
    data object OnAnimationComplete : SplashEvent()
    data object OnImageClicked : SplashEvent()
    data object OnSplashScreenOpened : SplashEvent()
}

sealed class SplashAction {
    data object NavigateToMain : SplashAction()
}
