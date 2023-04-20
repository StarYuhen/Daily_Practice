package com.yin.myapplication.database

import android.content.Context
import androidx.room.Room
import androidx.room.RoomDatabase





// 定义数据库初始化
@androidx.room.Database(entities = [UserTable::class], version = 1)
abstract class UserTableDatabase : RoomDatabase() {
    // 定义数据库操作接口
    abstract fun user(): UserData

    companion object{
        @Volatile
        private var db:UserTableDatabase? = null;
        // 获取DB
        fun getDB(context: Context):UserTableDatabase?{
                synchronized(UserTableDatabase::class.java){
                    if(db==null){
                        db=Room.databaseBuilder(context.applicationContext,UserTableDatabase::class.java,"user").
                        fallbackToDestructiveMigration().
                        build()
                    }
            }
            return db
        }
    }
}