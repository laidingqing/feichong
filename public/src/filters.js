'use strict';
define(['angular'], function(angular) {
  /* Services */
  angular.module('app.filters', [])
  .filter('BusinessStatus', function() {
    return function(input) {
      switch (input) {
        case 0:
          return "未开始";
          break;
        case 1:
          return "取单";
          break;
        case 2:
          return "做账";
          break;
        case 3:
          return "报税";
          break;
        case 4:
          return "回访";
          break;
        case 5:
          return "完成";
          break;
        default:
          return "未执行";
      }

    };
  });
});
