# Import flow

使用須知: 路徑固定不可變更，若變更須調整logs/logs.go和web/ws/client/client.go中的config import路徑，且須修改此文件敘述

PS:若已加入過的專案跳至第6步，初始化submodule 內容

## 1. 新專案使用時請從master切一分支並命名為你的專案名稱並改用小寫+底線組合命名
    Ex: 專案 gameServer
    $ git checkout -b game_server

## 2. 將logs/logs.go 和 web/ws/client/client.go 和 caches 底下所有.go中使用到的config import路徑前贅調整成golang module的名稱，新專案module名稱請查案該專案的go.mod第一行module
    Ex: module game_server
    將 "mgmt/common/web/response" 改成 "game_server/common/web/response"

## 3. 將子專案新增至新專案中 
    $ git submodule add http://gltw.6633663.com/Backend_GSGame/common.git common
   若該專案曾經使用過但後來移除可使用底下cmd重新加回
    $ git submodule add -f http://gltw.6633663.com/Backend_GSGame/common.git common

## 4. cd 至"專案"最外層資料夾中
    Ex: 後端的gameServer專案中的 gameserver層，不可進入wow_gaming or pkg 資料夾內下
    $ cd gameserver

## 5. 變更submodule branch
    $ git config -f .gitmodules submodule.common.branch xxxxx(branch)
    Ex:
    $ git config -f .gitmodules submodule.common.branch game_server

## 6. 初始化子專案
    $ git submodule init

## 7. 若該子專案資料夾內無資料使用底下cmd下載
    $ git submodule update --init

## 8. 若還是無法則使用底下cmd，再無法請自行查詢網路
    $ git submodule update --remote --recursive

## 9. go run 執行新的專案
    $ go run main.go