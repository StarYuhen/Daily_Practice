package com.yin.myapplication.database

import androidx.room.ColumnInfo
import androidx.room.Dao
import androidx.room.Entity
import androidx.room.Insert
import androidx.room.PrimaryKey
import androidx.room.Query
import io.reactivex.annotations.NonNull
import org.jetbrains.annotations.NonNls


// 使用data关键字
@Entity(tableName = "user")
data class UserTable (
    // 用于本地数据库操作
    @PrimaryKey @NonNls @ColumnInfo("username") val username: String,
    @ColumnInfo("password")  @NonNls  val password: String,
    @ColumnInfo("email")  @NonNls  val email: String,
    // 用户token，由username和password生成，使用ws连接后直接使用token，不允许使用username和password
   @ColumnInfo("token") val token:String
)


// 创建dao，实现接口函数
@Dao
interface UserData{

    // 查询token
    @Query("SELECT token FROM user")
    fun GetToken():String
    // 写入用户
    @Insert
    fun UserInsert(userTable: UserTable):Long

    // 获取用户所有信息
    @Query("SELECT * FROM user where username=:username")
    fun GetUser(username: String?):UserTable

}