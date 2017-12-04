/*global define */

'use strict';

define(function() {
  var controllers = {};

  // 登录控制器
  controllers.HomeCtrl = function($scope, $rootScope, $location) {
    $scope.login = function() {
      $location.path("/dashboard")
    }
  }
  controllers.HomeCtrl.$inject = ['$scope', '$rootScope', '$location'];

  // 面板控制器
  controllers.DashboardCtrl = function($scope, $rootScope) {
    $scope.title = "面板"
  }
  controllers.DashboardCtrl.$inject = ['$scope', '$rootScope'];

  // 订单控制器
  controllers.OrdersCtrl = function($scope, $rootScope, $uibModal, OrderService) {
    $scope.currentPageNum = 6;
    $scope.totalItems = 0;
    $scope.count = 10;
    $scope.currentPage = 1;
    $scope.maxSize = 8;
    $scope.pageSize = 10;
    $scope.param = {};
    $scope.data = {}; // 参数
    $scope.data.param = $scope.param;
    $scope.pagination = {};
    $scope.pagination.data = [];
    $scope.queryList = function(page) {
      $scope.data.startIndex = (page - 1) * $scope.pageSize;
      $scope.data.size = $scope.size;
      OrderService.getOrders($scope.data.startIndex, $scope.data.size, 1, function(res) {
        $scope.pagination.data = res.data.data;
        $scope.totalItems = res.totalCount;
        console.log(res.data)
      }, function(err) {
        console.log(err)
      })
    }
    $scope.showFormModal = function () {
        var modalInstance = $uibModal.open({
            templateUrl: '../components/orderFormModal.html?6',
            controller: controllers.NewOrderCtrl,
            size: 'lg',
            resolve: {
            }
        });
        modalInstance.result.then(function (newOutputData) {
          console.log(newOutputData)
            if(newOutputData){
              $scope.queryList(1)
            }
        });
    }
    $scope.queryList($scope.currentPage)
  }
  controllers.OrdersCtrl.$inject = ['$scope', '$rootScope', '$uibModal', 'OrderService'];

  // 业务跟踪控制器
  controllers.TracksCtrl = function($scope, $rootScope, $state, OrderService) {

    $scope.totalItems = 64;
    $scope.currentPage = 1;
    $scope.itemsPerPage = 10;

    $scope.pagination = {};
    $scope.pagination.data = [];

    $scope.$watch("currentPage", function() {
      queryList($scope.currentPage);
    });

    var queryList = function(page) {
      var page = (page - 1) * $scope.itemsPerPage;
      var size = $scope.itemsPerPage;
      OrderService.getOrders(page, size, 1, function(res) {
        $scope.pagination.data = res.data.data;
        $scope.totalItems = res.totalCount;
        console.log(res.data)
      }, function(err) {
        console.log(err)
      })
    }

    $scope.trackDetail = function(orderId, orderNO){
      console.log(orderId, $state)
      $state.go('dashboard.business', {id: orderId, orderNO: orderNO });
    }

  }
  controllers.TracksCtrl.$inject = ['$scope', '$rootScope', '$state', 'OrderService'];

  controllers.NewOrderCtrl = function($scope, $uibModalInstance, $filter, OrderService, UserService) {
    $scope.users = []
    $scope.companies = []
    $scope.order = {
      orderNO: $filter('date')(new Date(), 'yyyyMMddHHmmss')
    }
    $scope.selected
    var queryList = function() {
      UserService.getUserBySelf(function(response){
        $scope.users = response.data
        console.log($scope.users)
      }, function(err){
        console.log(err)
      })
      UserService.getUserByEnterprise(function(res){
        $scope.companies = res.data
        console.log(res.data)
      }, function(err){
        console.log(err)
      })
    }
    queryList();

    $scope.ok = function(){
        console.log($scope.order)
        OrderService.postOrder($scope.order, function(res){
          console.log(res)
          $uibModalInstance.close(true)
        }, function(err){
          console.log(err)
        })
    }
  }

  controllers.NewOrderCtrl.$inject = ['$scope', '$uibModalInstance', '$filter', 'OrderService', 'UserService'];

  // 客户管理控制器
  controllers.CustomersCtrl = function($scope, $rootScope, $uibModal, UserService) {
    $scope.totalItems = 64;
    $scope.currentPage = 1;
    $scope.itemsPerPage = 10;

    $scope.pagination = {};
    $scope.pagination.data = [];

    $scope.$watch("currentPage", function() {
      queryList($scope.currentPage);
    });

    var queryList = function(page) {
      var page = (page - 1) * $scope.itemsPerPage;
      var size = $scope.itemsPerPage;
      UserService.getUsers(page, size, function(res) {
        $scope.pagination.data = res.data.data;
        $scope.totalItems = res.totalCount;
        console.log(res.data)
      }, function(err) {
        console.log(err)
      })
    }
    $scope.upateProfile = function(id){
      var modalInstance = $uibModal.open({
          templateUrl: '../components/userFormModal.html?7',
          controller: controllers.UpdateProfileCtrl,
          size: 'lg',
          resolve: {
            updateUserId: function(){
              return id
            }
          }
      });
      modalInstance.result.then(function (newOutputData) {
        console.log(newOutputData)
          if(newOutputData){
            queryList(1)
          }
      });
    }
    //更新登录信息
    $scope.upateSession = function(id){
      var modalInstance = $uibModal.open({
          templateUrl: '../components/securityModal.html?1',
          controller: function($scope, $uibModalInstance, UserService, updateUserId){
              $scope.user = {}
              $scope.error = false
              console.log("更新登录会话信息", updateUserId)
              UserService.getUser(updateUserId, function(res){
                $scope.user = res
              }, function(err){
                $scope.error = true
                $scope.errorText = err
              })

              $scope.submit = function(){
                UserService.checkUserName($scope.user.username, function(res){
                    console.log(res)
                    if( res.id != ""){
                      console.log("已存在用户")
                      $scope.error = true
                      $scope.errText = "已存在相同的登录用户名，请重新输入！"
                    }else{
                      UserService.putUserSecurity($scope.user, function(data){
                        console.log(data)v     ft
                      })
                    }
                }, function(err){
                  console.log(err)
                })
              }
          },
          size: 'lg',
          resolve: {
            updateUserId: function(){
              return id
            }
          }
      });
      modalInstance.result.then(function (newOutputData) {
        console.log(newOutputData)
          if(newOutputData){
            queryList(1)
          }
      });
    }
  }
  controllers.CustomersCtrl.$inject = ['$scope', '$rootScope', '$uibModal', 'UserService'];

  controllers.UpdateProfileCtrl = function($scope, $uibModalInstance, UserService, updateUserId){
    $scope.user = {}
    UserService.getUser(updateUserId, function(res){
      console.log(res)
      $scope.user = res
    }, function(err){
      console.log(err)
    })

    $scope.submit = function () {
      console.log($scope.user)
      if( $scope.user.username != ""){
        //check username
      }

      UserService.putUser($scope.user, function(res){
        $uibModalInstance.close(true);
      })

    };
    $scope.cancel = function(){
      $uibModalInstance.close();
    }
  }

  controllers.UpdateProfileCtrl.$inject = ['$scope', '$uibModalInstance', 'UserService', 'updateUserId'];

  controllers.BusinessCtrl = function($scope, $uibModal, $stateParams, UserService, OrderService, BusinessService){
    console.log($stateParams, $stateParams.id)
    $scope.business = []
    $scope.orderNO = $stateParams.orderNO
    var init = function(){
      OrderService.getBusinessByOrder($stateParams.id, function(data){
        $scope.business = data || []
        console.log(data)
      }, function(err){
        console.log(err)
      })
    }

    init()

    $scope.showCapital = function(orderId, businessId){
      var modalInstance = $uibModal.open({
          templateUrl: '../components/capitalInfoModal.html?2',
          controller: function($scope, $uibModalInstance, OrderService){
            $scope.err = false
            $scope.errorText = ""
            $scope.capitalInfo = {}
            OrderService.getBusinessByID(orderId, businessId, function(data){
              $scope.capitalInfo = data.capitalInfo
              console.log(data.capitalInfo)
            }, function(err){
              console.log(err)
            })
            $scope.cancel = function(){
                $uibModalInstance.close();
            }
            $scope.ok = function(){
              BusinessService.putCapitalInfo(businessId, $scope.capitalInfo, function(res){
                if( res.status == 200){
                  $scope.capitalInfo = data
                  console.log($scope.capitalInfo, data)
                  $uibModalInstance.close();
                }
              }, function(err){
                $scope.err = true
                $scope.errorText = err.message
              })
            }
          },
          size: 'lg',
          resolve: {

          }
      });

      modalInstance.result.then(function(){

      }, function(res){

      })
    }
    $scope.showProfit = function(orderId, businessId){
      var modalInstance = $uibModal.open({
          templateUrl: '../components/profitInfoModal.html?1',
          controller: function($scope, $uibModalInstance, OrderService){
            $scope.profitInfo = {}
            $scope.err = false
            $scope.errorText = ""
            OrderService.getBusinessByID(orderId, businessId, function(data){
              $scope.profitInfo = data.profitInfo
              console.log(data.profitInfo)
            }, function(err){
              console.log(err)
            })

            $scope.cancel = function(){
                $uibModalInstance.close();
            }
            $scope.ok = function(){
              BusinessService.putProfitInfo(businessId, $scope.profitInfo, function(res){
                if(res.status == 200){
                  $scope.profitInfo = res.data
                  console.log(data.profitInfo)
                  $uibModalInstance.close();
                }
              }, function(err){
                $scope.err = true
                $scope.errorText = err.message
              })
            }

          },
          size: 'lg',
          resolve: {

          }
      });
      modalInstance.result.then(function(){

      }, function(res){

      })
    }
    $scope.showTax = function(orderId, businessId){
      var modalInstance = $uibModal.open({
          templateUrl: '../components/taxInfoModal.html?1',
          controller: function($scope, $uibModalInstance, OrderService){
            $scope.err = false
            $scope.errorText = ""
            $scope.taxInfo = {}
            OrderService.getBusinessByID(orderId, businessId, function(data){
              $scope.taxInfo = data.taxInfo
              console.log(data.taxInfo)
            }, function(err){
              $scope.err = true
              $scope.errorText = err.message
            })

            $scope.cancel = function(){
                $uibModalInstance.close();
            }
            $scope.ok = function(){
              BusinessService.putTaxInfo(businessId, $scope.taxInfo, function(res){
                if( res.status == 200){
                  $scope.taxInfo = data.taxInfo
                  console.log($scope.taxInfo, res)
                  $uibModalInstance.close();
                }
              }, function(err){
                console.log(err)
              })
            }
          },
          size: 'lg',
          resolve: {

          }
      });
      modalInstance.result.then(function(){

      }, function(res){

      })
    }
  }

  controllers.BusinessCtrl.$inject = ['$scope', '$uibModal', '$stateParams', 'UserService', 'OrderService', 'BusinessService'];

  controllers.FeedbackCtrl = function($scope, $uibModal, $stateParams, UserService, OrderService, BusinessService){

  }
  controllers.BusinessCtrl.$inject = ['$scope', '$uibModal', '$stateParams', 'UserService', 'OrderService', 'BusinessService'];

  return controllers;
});
