# Goroutine-Tutorial
This is an example for tutorial of goroutine.
Please reference my youtube channel's video or blog's articles.

## Context
1. + video: [Golang 教學系列 - 何謂Groutine? Goroutine 簡單教學! | 肯尼攻城獅](https://www.youtube.com/watch?v=nEWcQmJgP-0&list=PLIud7iV0oWk-oQ6Da7WyrGShZ-auYZq8e&index=1)
   + article: [Golang 教學系列 - 何謂 Goroutine](https://blog.kennycoder.io/2020/12/13/Golang%E6%95%99%E5%AD%B8%E7%B3%BB%E5%88%97-%E4%BD%95%E8%AC%82Goroutine/)
2. + video: [Golang 教學系列 - 何謂WaitGroup? 等待Goroutine的好幫手! | 肯尼攻城獅](https://www.youtube.com/watch?v=1-Z4XCCU3B4&list=PLIud7iV0oWk-oQ6Da7WyrGShZ-auYZq8e&index=2)
   + article: [Golang 教學系列 - 何謂 WaitGroup? 等待 Goroutine 的好幫手！](https://blog.kennycoder.io/2020/12/18/Golang%E6%95%99%E5%AD%B8%E7%B3%BB%E5%88%97-%E4%BD%95%E8%AC%82WaitGroup-%E7%AD%89%E5%BE%85Goroutine%E7%9A%84%E5%A5%BD%E5%B9%AB%E6%89%8B/)
3. + video: [Golang 教學系列 - WaitGroup 常見的坑以及應用介紹! | 肯尼攻城獅](https://www.youtube.com/watch?v=5Yes7NTgShk&list=PLIud7iV0oWk-oQ6Da7WyrGShZ-auYZq8e&index=3)
   + article: [Golang 教學系列 - WaitGroup 常見的坑以及應用介紹!](https://blog.kennycoder.io/2020/12/20/Golang%E6%95%99%E5%AD%B8%E7%B3%BB%E5%88%97-WaitGroup%E5%B8%B8%E8%A6%8B%E7%9A%84%E5%9D%91%E4%BB%A5%E5%8F%8A%E6%87%89%E7%94%A8%E4%BB%8B%E7%B4%B9/)
4. + video: [Golang 教學系列 - 何謂Channel? 先從宣告Channel開始學起! | 肯尼攻城獅](https://www.youtube.com/watch?v=GFsaafzG7XU&list=PLIud7iV0oWk-oQ6Da7WyrGShZ-auYZq8e&index=4)
   + article: [Golang 教學系列 - 何謂Channel? 先從宣告Channel開始學起!](https://blog.kennycoder.io/2020/12/23/Golang%E6%95%99%E5%AD%B8%E7%B3%BB%E5%88%97-%E4%BD%95%E8%AC%82Channel-%E5%85%88%E5%BE%9E%E5%AE%A3%E5%91%8AChannel%E9%96%8B%E5%A7%8B%E5%AD%B8%E8%B5%B7/)
5. + video: [Golang 教學系列 - channel裡面資料什麼都能放? 單雙向channel的差別? | 肯尼攻城獅](https://www.youtube.com/watch?v=H1wOqwQetUM&list=PLIud7iV0oWk-oQ6Da7WyrGShZ-auYZq8e&index=5)
   + article: [Golang 教學系列 - channel裡面資料什麼都能放? 單雙向channel的差別?](https://blog.kennycoder.io/2020/12/29/Golang%E6%95%99%E5%AD%B8%E7%B3%BB%E5%88%97-channel%E8%A3%A1%E9%9D%A2%E8%B3%87%E6%96%99%E4%BB%80%E9%BA%BC%E9%83%BD%E8%83%BD%E6%94%BE-%E5%96%AE%E9%9B%99%E5%90%91channel%E7%9A%84%E5%B7%AE%E5%88%A5/)
6. + video: [Golang 教學系列 - close 與 for range 搭配channel觀念教學! | 肯尼攻城獅](https://www.youtube.com/watch?v=9HtkKad33lM&list=PLIud7iV0oWk-oQ6Da7WyrGShZ-auYZq8e&index=6)
7. + video: [Golang 教學系列 - channel的應用範例示範! 讓你更懂得channel使用! | 肯尼攻城獅](https://www.youtube.com/watch?v=ToDAN8IndMI&list=PLIud7iV0oWk-oQ6Da7WyrGShZ-auYZq8e&index=7)
8. + video: [Golang 教學系列 - select 與 channel的搭配教學! | 肯尼攻城獅](https://www.youtube.com/watch?v=FN0XMV4q5-A&list=PLIud7iV0oWk-oQ6Da7WyrGShZ-auYZq8e&index=8)
9. + video: [Golang 教學系列 - timer與channel教學示範 實現定時器功能! | 肯尼攻城獅](https://www.youtube.com/watch?v=Rs9-0m3DsBE&list=PLIud7iV0oWk-oQ6Da7WyrGShZ-auYZq8e&index=9)
10. + video: [Golang 教學系列 - time after vs afterFunc的差別? ticker教學示範~| 肯尼攻城獅](https://www.youtube.com/watch?v=ZoEaiK-7orI&list=PLIud7iV0oWk-oQ6Da7WyrGShZ-auYZq8e&index=10)
11. + video: [Golang 教學系列 - cronJob 陽春功能用ticker實現示範! | 肯尼攻城獅](https://www.youtube.com/watch?v=akS8ZiDOit4&list=PLIud7iV0oWk-oQ6Da7WyrGShZ-auYZq8e&index=11)
12. + video: [Golang 教學系列 - 從讀原始碼講解context的運作方式](https://www.youtube.com/watch?v=ix9IMCsemTw&list=PLIud7iV0oWk-oQ6Da7WyrGShZ-auYZq8e&index=12)
13. + video: [Golang 教學系列 - 可以cancel的context是什麼？如何cancel？](https://www.youtube.com/watch?v=RM29P-VkWrc&list=PLIud7iV0oWk-oQ6Da7WyrGShZ-auYZq8e&index=13)
14. + video: [Golang 教學系列 - 再來一個cancel context的介紹！](https://www.youtube.com/watch?v=FddDt3JgMAo&list=PLIud7iV0oWk-oQ6Da7WyrGShZ-auYZq8e&index=14)
15. + video: [Golang 教學系列 - 結合多種不同類型的context的例子示範！](https://www.youtube.com/watch?v=uwyG6msSOSY&list=PLIud7iV0oWk-oQ6Da7WyrGShZ-auYZq8e&index=15)
16. + video: [Golang 教學系列 - context還可以定時取消？還可以傳值？| 肯尼攻城獅](https://www.youtube.com/watch?v=H-z8kOBB_MA&list=PLIud7iV0oWk-oQ6Da7WyrGShZ-auYZq8e&index=16)
<p>to be continued...</p>