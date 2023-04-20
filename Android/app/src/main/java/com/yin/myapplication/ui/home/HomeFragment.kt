package com.yin.myapplication.ui.home

import android.os.Bundle
import android.view.LayoutInflater
import android.view.View
import android.view.ViewGroup
import android.widget.TextView
import android.widget.Toast
import androidx.fragment.app.Fragment
import androidx.lifecycle.ViewModelProvider
import androidx.lifecycle.setViewTreeLifecycleOwner
import androidx.navigation.findNavController
import com.yin.myapplication.R
import com.yin.myapplication.database.UserTableDatabase
import com.yin.myapplication.databinding.FragmentHomeBinding

class HomeFragment : Fragment() {

    private var _binding: FragmentHomeBinding? = null

    // This property is only valid between onCreateView and
    // onDestroyView.
    private val binding get() = _binding!!

    override fun onCreateView(
        inflater: LayoutInflater,
        container: ViewGroup?,
        savedInstanceState: Bundle?
    ): View {
        val homeViewModel =
            ViewModelProvider(this).get(HomeViewModel::class.java)

        _binding = FragmentHomeBinding.inflate(inflater, container, false)
        // 不能为空，获取application
//        val application = requireNotNull(this.activity).application
//        // 绑定数据库
//        val dataSource = UserTableDatabase.getDB(application)?.user()
//
//        // 获得内容
//        val viewUserModel= ViewModelProvider(this, HomeViewModelFactory(dataSource!!, application)).
//        get(HomeViewModel::class.java)
//        // 绑定最新的方法，且引入dataBing依赖包
//        binding.viewUserModel= viewUserModel
//        binding.lifecycleOwner = this

        // 绑定视图
        val root: View = binding.root
        // 获取绑定id
        val textView: TextView = binding.textHome
        // 使用ViewModel中的数据进行text映射，一旦有改变就映射
        homeViewModel.text.observe(viewLifecycleOwner) {
            textView.text = it
        }


        return root
    }

    // 一旦当视图被销毁，就自动销毁绑定的视图 FragmentHomeBinding
    override fun onDestroyView() {
        super.onDestroyView()
        _binding = null
    }


    override fun onViewCreated(view: View, savedInstanceState: Bundle?) {
        super.onViewCreated(view, savedInstanceState)
        // 通过id获取视图
        val viewButtonGOAction=view.findViewById<View>(R.id.gotoAction)
        // 设置点击事件
        // 创建前往意图的方法
        viewButtonGOAction?.setOnClickListener{
            Toast.makeText(context, "前往控制路由", Toast.LENGTH_SHORT).show()

        }
    }
}