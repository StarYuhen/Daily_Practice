package com.yin.myapplication.ui.gallery

import android.app.NativeActivity
import android.content.ContentValues.TAG
import android.os.Bundle
import android.util.Log
import android.view.LayoutInflater
import android.view.View
import android.view.ViewGroup
import android.widget.ImageView
import android.widget.TextView
import android.widget.Toast
import androidx.drawerlayout.widget.DrawerLayout
import androidx.fragment.app.Fragment
import androidx.lifecycle.ViewModelProvider
import androidx.navigation.findNavController
import androidx.navigation.ui.AppBarConfiguration
import com.yin.myapplication.R
import com.yin.myapplication.databinding.FragmentGalleryBinding
import kotlin.math.log
import kotlin.random.Random

class GalleryFragment : Fragment(){

    private var _binding: FragmentGalleryBinding? = null

    // This property is only valid between onCreateView and
    // onDestroyView.
    private val binding get() = _binding!!

    override fun onCreateView(
        inflater: LayoutInflater,
        container: ViewGroup?,
        savedInstanceState: Bundle?
    ): View {
        val galleryViewModel =
            ViewModelProvider(this).get(GalleryViewModel::class.java)

        _binding = FragmentGalleryBinding.inflate(inflater, container, false)
        val root: View = binding.root

        val textView: TextView = binding.textGallery
        galleryViewModel.text.observe(viewLifecycleOwner) {
            textView.text = it
        }



            return root
        }

    override fun onDestroyView() {
        super.onDestroyView()
        _binding = null
    }





    /*
    onViewCreated()方法是Fragment生命周期中的一个回调方法，它在Fragment的视图层次结构被创建后立即调用。具体来说，当Fragment的根View和所有子View都被实例化和添加到视图层次结构中之后，这个方法就会被触发。
    onViewCreated()方法的作用是允许你对Fragment的视图层次结构进行初始化或者其他相关操作。在这个方法中，你可以通过view参数获取到Fragment的根View对象，并对其进行相关操作，例如设置点击事件监听器、绑定数据等。
    在实际开发中，你通常会在onCreateView()方法中返回Fragment的根View对象，并在onViewCreated()方法中为其中的子View设置一些操作。例如，你可以在onCreateView()方法中使用LayoutInflater对象来将一个布局文件转换成一个View对象，然后在onViewCreated()方法中使用findViewById()方法来找到其中的一些子View并为它们设置操作。
     */
    override fun onViewCreated(view: View, savedInstanceState: Bundle?) {
        super.onViewCreated(view, savedInstanceState)
        val viewButtonGoTOHome = view.findViewById<View>(R.id.buttonGoToHome)
        val viewImg=view.findViewById<ImageView>(R.id.Dice)



        viewButtonGoTOHome?.setOnClickListener {
            Toast.makeText(context, "点击了按钮", Toast.LENGTH_SHORT).show()
            // 同时更换字体，使用绑定后的布局Binding中，使用textGallery(是控件名称)进行字体更换
            when(Random.Default.nextInt(6)+1){
                1->viewImg.setImageResource(R.drawable.dice_1)
                2->viewImg.setImageResource(R.drawable.dice_2)
                3->viewImg.setImageResource(R.drawable.dice_3)
                4->viewImg.setImageResource(R.drawable.dice_4)
                5->viewImg.setImageResource(R.drawable.dice_5)
                6->viewImg.setImageResource(R.drawable.dice_6)
            }

            binding.textGallery.text="你好"
        }




    }




}