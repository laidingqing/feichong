/*global define */

'use strict';

define(function() {
    var controllers = {};

    controllers.HomeCtrl = function($scope, $rootScope, $location) {
        $scope.login = function() {
          $location.path("/dashboard")
        }
    }
    controllers.HomeCtrl.$inject = ['$scope', '$rootScope', '$location'];

    controllers.DashboardCtrl = function($scope, $rootScope) {
      $scope.title = "面板"
    }
    controllers.DashboardCtrl.$inject = ['$scope', '$rootScope'];

    controllers.CustomersCtrl = function($scope, $rootScope) {
      $scope.title = "客户管理"
    }
    controllers.CustomersCtrl.$inject = ['$scope', '$rootScope'];


    return controllers;
});
