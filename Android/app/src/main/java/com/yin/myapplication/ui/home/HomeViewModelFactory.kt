package com.yin.myapplication.ui.home

import android.app.Application
import androidx.lifecycle.ViewModel
import androidx.lifecycle.ViewModelProvider
import androidx.lifecycle.viewmodel.CreationExtras
import com.yin.myapplication.database.UserData
import com.yin.myapplication.database.UserTable


// 扩展Factory
//class HomeViewModelFactory(
//    private val userData: UserData,
//    private val application: Application
//) : ViewModelProvider.Factory {
//    @Suppress("unchecked_cast")
//    override fun <T : ViewModel> create(modelClass: Class<T>, extras: CreationExtras): T {
//        if (modelClass.isAssignableFrom(HomeViewModel::class.java)) {
//            return HomeViewModel(userData, application) as T
//        }
//        throw IllegalArgumentException("Unknown ViewModel class")
//    }
//}
