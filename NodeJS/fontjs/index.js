import { fontSplit } from "@konghayao/cn-font-split";

fontSplit({
    FontPath: "./fonts/LXGWWenKaiGBScreen.ttf",
    destFold: "./build",
    css: {
        fontFamily: "LXGWWenKaiGBScreen", // 不建议使用，我们已经有内置的解析模块了
        // fontWeight: 400,
    },
    targetType: "woff2", // ttf woff woff2；注意 eot 文件在浏览器中的支持度非常低，所以不进行支持
    chunkSize: 50 * 1024, // 分割大小
    testHTML: true, // 输出一份 html 报告文件
    reporter: true, // 输出 json 报告
});

/*

// 引用自定义字体分割文件
* {
    font-family:LXGWWenKaiGBScreen;
}
body {
    font-family: LXGWWenKaiGBScreen;
}


// 设置字体储存
  // 发起请求获取字体文件
  fetch('https://api.yuhenm.com/api/file/source?name=LXGWWenKaiGBScreen.woff', {
    cache: 'force-cache' // 强制缓存
  })
    .then(response => response.blob())
    .then(blob => {
      // 创建字体文件的临时URL
      const fontUrl = URL.createObjectURL(blob);

      // 创建<style>标签
      const style = document.createElement('style');
      // 设置<style>标签的内容
      style.textContent = `
        @font-face {
          font-family: LXGWWenKaiGBScreen;
          font-size: 16px;
          font-style: normal;
          font-display: swap;
          src: url(${fontUrl}) format('woff');
        }
        * {
          font-family: LXGWWenKaiGBScreen;
        }
        body {
          font-family: LXGWWenKaiGBScreen;
        }
      `;
      // 将<style>标签添加到<head>中
      document.head.appendChild(style);

      // 将字体缓存状态标记为已加载
      localStorage.setItem('customFontLoaded', 'true');
    });

    // 强制缓存css文件
const link = document.createElement('link');
link.rel = 'stylesheet';
link.href = 'https://api.yuhenm.com/api/file/source?name=result.css';
link.setAttribute('as', 'style');
link.setAttribute('integrity', 'fe7b7acdc41c0a1fa26663923d0401f6g6676k65y7v'); // 替换为您的样式表的哈希值

// 使用 fetch API 进行预加载和缓存
fetch(link.href, { cache: 'force-cache' })
  .then(response => response.text())
  .then(css => {
    link.textContent = css; // 将样式表内容赋值给 link 元素

    // 将 link 元素添加到文档头部
    document.head.appendChild(link);
  })
  .catch(error => {
    // 处理加载样式表的错误
    console.error('Failed to load stylesheet:', error);
  });


 */