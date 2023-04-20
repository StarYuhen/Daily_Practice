package com.yin.myapplication

import android.util.Log
import androidx.room.Room
import androidx.room.Room.inMemoryDatabaseBuilder
import androidx.test.ext.junit.runners.AndroidJUnit4
import androidx.test.platform.app.InstrumentationRegistry
import com.yin.myapplication.database.UserData
import com.yin.myapplication.database.UserTable
import com.yin.myapplication.database.UserTableDatabase
import junit.framework.TestCase.assertEquals
import org.junit.Before
import org.junit.Test
import org.junit.runner.RunWith
import java.io.IOException
import java.lang.StrictMath.log


@RunWith(AndroidJUnit4::class)
class UserTableTest {
    private lateinit var userTableData: UserData
    private lateinit var userTableDatabase: UserTableDatabase

    @Before
    fun createDb() {
        // 创建数据库
        val context = InstrumentationRegistry.getInstrumentation().targetContext
        userTableDatabase=inMemoryDatabaseBuilder(context,UserTableDatabase::class.java).
            allowMainThreadQueries().
            build()
        userTableData=userTableDatabase.user()
    }

    // 关闭数据库
    @Throws(IOException::class)
    fun closeDb(){
        userTableDatabase.close()
    }

    // 测试数据库
    @Test
    @Throws(Exception::class)
    fun testUserTable(){
        // 测试插入
        userTableData.UserInsert(UserTable("yin","123456","",""))
        // 测试查询
        val user=userTableData.GetUser("yin")
        // 测试结果
        assertEquals(user.password,"123456")
        Log.d("user查询内容",user.username)
    }

}