package com.yin.myapplication

import android.os.Bundle
import android.service.autofill.Sanitizer
import android.view.Menu
import com.google.android.material.snackbar.Snackbar
import com.google.android.material.navigation.NavigationView
import androidx.navigation.findNavController
import androidx.navigation.ui.AppBarConfiguration
import androidx.navigation.ui.navigateUp
import androidx.navigation.ui.setupActionBarWithNavController
import androidx.navigation.ui.setupWithNavController
import androidx.drawerlayout.widget.DrawerLayout
import androidx.appcompat.app.AppCompatActivity
import com.yin.myapplication.databinding.ActivityMainBinding

class MainActivity : AppCompatActivity() {

    private lateinit var appBarConfiguration: AppBarConfiguration
    private lateinit var binding: ActivityMainBinding

    override fun onCreate(savedInstanceState: Bundle?) {
        super.onCreate(savedInstanceState)


        // 将主视图文件进行绑定
        binding = ActivityMainBinding.inflate(layoutInflater)
        setContentView(binding.root)

        // 操作导航栏
        setSupportActionBar(binding.appBarMain.toolbar)


        // 用于绑定一个浮动按钮
        binding.appBarMain.fab.setOnClickListener { view ->
            Snackbar.make(view,"启动成功",Snackbar.LENGTH_LONG)
                .setAction("Action",null).show()
        }
        // 抽屉式布局
        val drawerLayout: DrawerLayout = binding.drawerLayout
        // 导航栏目
        val navView: NavigationView = binding.navView
        // 导航控制器，实现导航的空值功能
        val navController = findNavController(R.id.nav_host_fragment_content_main)
        // Passing each menu ID as a set of Ids because each
        // menu should be considered as top level destinations.
        appBarConfiguration = AppBarConfiguration(
            // 实现导航功能     <fragment 这是布局导航，选择xml文件的id后进行视图切换
            setOf(
                R.id.nav_home, R.id.nav_gallery, R.id.nav_slideshow
            ), drawerLayout
        )
        // 绑定导航栏目
        setupActionBarWithNavController(navController, appBarConfiguration)
        //      val navController = findNavController(R.id.nav_host_fragment_content_main) 与他进行绑定，而后切换控制器刷新
        navView.setupWithNavController(navController)
    }


    // 创建选项菜单
    /*
        <item
        android:id="@+id/action_settings"
        android:orderInCategory="100"
        android:title="@string/action_settings"
        app:showAsAction="never" />
     */
    override fun onCreateOptionsMenu(menu: Menu): Boolean {
        // Inflate the menu; this adds items to the action bar if it is present.
        menuInflater.inflate(R.menu.main, menu)
        return true
    }

    // 用于绑定用的的返回操作，会自动返回上一级界面，绑定的ID nav_host_fragment_content_main界面
    override fun onSupportNavigateUp(): Boolean {
        val navController = findNavController(R.id.nav_host_fragment_content_main)
        return navController.navigateUp(appBarConfiguration) || super.onSupportNavigateUp()
    }
}