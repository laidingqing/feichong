'use strict';
define(['angular'], function(angular) {
  /* Services */
  angular.module('app.directive', [])
  .directive('selectDate',[function(){
        return{
            restrict:'AE',
            replace:true,
            scope:{
                startYear:'@',
                endYear:'@',
                ngModel:'='
            },
            templateUrl:'../partials/inputEle.html',
            controller:['$scope',function($scope){
                var fromYear=$scope.startYear?parseInt($scope.startYear):(new Date()).getFullYear(),
                    toYear=$scope.endYear?parseInt($scope.endYear):(new Date()).getFullYear()+15,
                    yearArr=[];
                for(var i=fromYear;i<=toYear;i++){
                    yearArr.push(i);
                }
                $scope.yearArr=yearArr;
                $scope.selectYear=$scope.yearArr[0];

                $scope.monthArr=[1,2,3,4,5,6,7,8,9,10,11,12];
                $scope.selectMonth=$scope.monthArr[0];
            }],
            link: function(scope, elem, attrs) {
                var newDate={
                    selectMonth: scope.selectMonth,
                    selectYear: scope.selectYear
                };
                scope.ngModel=newDate;
                scope.$watch('selectYear',function(newVal,oldVal){
                    if(newVal&&newVal!==oldVal){
                        newDate.selectYear=newVal;
                        scope.ngModel=newDate;
                    }

                });
                scope.$watch('selectMonth',function(newVal,oldVal){
                    if(newVal&&newVal!==oldVal){
                        newDate.selectMonth=newVal;
                        scope.ngModel=newDate;
                    }
                });
            }
        }
    }]);
})
