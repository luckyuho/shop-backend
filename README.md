# 微型電商

## 設計採前後端分離

[前端(vue)]() | [後端(golang)]()


## 整體流程

使用者可註冊登入，或透過第三方認證的快速登入，並選擇想要購買的商品，透過第三方藍新金流支付購買商品。


## 功能

**登入：**

  * 使用者註冊與登入系統
  * 串接 [github 第三方登入](https://docs.github.com/zh/apps/oauth-apps/building-oauth-apps/authorizing-oauth-apps)，合併至原有的使用者註冊登入系統
  * 採用 cookie 緩存 jwt
   
**金流：**

  * 將訂單資訊傳送給[藍新金流(測試)](https://cwww.newebpay.com/)，藍新金流會提供支付方式
  * 付款的結果會分別回傳給前端與後端

**購物：**

  * 點擊購買物品，並確認成功付款後，會更新物品購買狀態，並記錄擁有者


## 流程介紹

**創建的表單：**

  * 使用者帳密表: user_table
  * 商品表: product
  * 訂單表: order

**登入系統：**

  * 使用者註冊或者採用快速登入(未登入過)，都會記錄在"使用者帳密資訊"的表單中。
  * github 的第三方登入流程解說可參考[連結](https://ithelp.ithome.com.tw/articles/10291817)，由 github 回傳的資訊中，使用 github 提供的 id 作為使用者姓名(因為這在github帳號中是定值)
  * 登入時則查找資料庫是否有該使用者資訊，有的話回傳 jwt 給前端

**金流系統：**

  * 串接[藍新金流(測試)](https://cwww.newebpay.com/)，本篇串接時，版本為 2.0
  * 使用[官方提供的手冊-1.0.6.pdf]()
  * 將商品以及訂單資訊，按照藍新手冊中的加密方式打包給藍新金流
  * 藍新確認訂單無誤後，畫面會跳轉到藍新金流所提供的付款畫面
  * 確認使用者的付款資訊後，會通知前端與後端訂單是否成功付款

**購物系統**

  * 選擇喜歡的商品後點擊購買並前往付款
  * 此時在 order 表單中會建立訂單編號
  * 第三方支付會回傳訂單是否成功付款
  * 若收到成功付款資訊後，會更新 product 中的產品資訊


## 影片介紹

登入採用第三方登入方式 
<center>
  <img src="https://github.com/luckyuho/shop-backend/blob/main/img/shop.gif" height="500"> 
</center>


 * ps. 最後成功支付完的跳轉畫面為google錯誤畫面，是因為藍新支付結束的跳轉與通知頁都必須是https的，但因為只有單台電腦，因此透過ngrok只能開啟一個https的網址，因為後端必須通知，才能更新訂單資訊，因此只好捨棄前端畫面，跳轉頁就只好亂設一通，如果前後端都部署到相對應的網域中，就不會有這個問題。


藍新支付收據
<center>
  <img src="https://github.com/luckyuho/shop-backend/blob/main/img/email.png" height="600"> 
</center>


## TodoList

  - [ ] 新增電子發票資訊
  - [ ] 前端的架構可以再修改一下


## 參考資料
  * [第三方登入流程](https://ithelp.ithome.com.tw/articles/10291817): [https://ithelp.ithome.com.tw/articles/10291817](https://ithelp.ithome.com.tw/articles/10291817)
