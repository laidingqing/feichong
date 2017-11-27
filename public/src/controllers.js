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
            templateUrl: '../components/orderFormModal.html?5',
            controller: controllers.NewOrderCtrl,
            size: 'lg',
            resolve: {
            }
        });
        return modalInstance;
    }
    $scope.queryList($scope.currentPage)
  }
  controllers.OrdersCtrl.$inject = ['$scope', '$rootScope', '$uibModal', 'OrderService'];

  // 业务跟踪控制器
  controllers.TracksCtrl = function($scope, $rootScope, OrderService) {

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

  }
  controllers.TracksCtrl.$inject = ['$scope', '$rootScope', 'OrderService'];

  controllers.NewOrderCtrl = function($scope, $uibModal, OrderService, UserService) {
    $scope.users = []
    $scope.selected
    var queryList = function() {
      UserService.getUserBySelf(function(response){
        $scope.users = response.data
        console.log($scope.users)
      }, function(err){
        console.log(err)
      })
    }
    queryList();
  }

  controllers.NewOrderCtrl.$inject = ['$scope', '$uibModal', 'OrderService', 'UserService'];

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
          templateUrl: '../components/userFormModal.html?6',
          controller: controllers.UpdateProfileCtrl,
          size: 'lg',
          resolve: {
            updateUserId: function(){
              return id
            }
          }
      });
      return modalInstance;
    }
  }
  controllers.CustomersCtrl.$inject = ['$scope', '$rootScope', '$uibModal', 'UserService'];

  controllers.UpdateProfileCtrl = function($scope, $uibModal, UserService, updateUserId){
    $scope.user = {}
    UserService.getUser(updateUserId, function(res){
      $scope.user = res
    }, function(err){
      console.log(err)
    })

    $scope.ok = function () {
      $uibModal.close();
    };
  }

  controllers.UpdateProfileCtrl.$inject = ['$scope', '$uibModal', 'UserService', 'updateUserId'];


  return controllers;
});
