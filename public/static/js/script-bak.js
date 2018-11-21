/*

This is a javascript file for omninotesweb
============

Author:  Suraj patil
Updated: January 2015
keyCode: n-110
*/


$.fn.serializeObject = function() {
	var o = {};
	var a = this.serializeArray();
	$.each(a, function() {
		if (o[this.name]) {
			if (!o[this.name].push) {
				o[this.name] = [ o[this.name] ];
			}
				o[this.name].push(this.value || '');
			} else {
				o[this.name] = this.value || '';
			}
		});
		return o;
}
		
$(document).ready(function(){

  $('#myform').bootstrapValidator({
        message: 'This value is not valid',
        feedbackIcons: {
            valid: 'glyphicon glyphicon-ok',
            invalid: 'glyphicon glyphicon-remove',
            validating: 'glyphicon glyphicon-refresh'
        },
        fields: {
            key_fileld: {
                message: '主键字段输入无效',
                validators: {
                    notEmpty: {
                        message: '主键字段不能为空'
                    },
                    stringLength: {
                        min: 6,
                        max: 30,
                        message: '字段字段长度不对！'
                    },
                    
                }
            },
			time_filed: {
                message: '时间字段无效',
                validators: {
                    notEmpty: {
                        message: '时间字段不能为空！'
                    },
                    stringLength: {
                        min: 6,
                        max: 30,
                        message: '字段字段长度不对！'
                    },
                    
                }
            },
		
        }
    }) 
		

});
