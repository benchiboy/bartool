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
		
            title: {
                message: '标题不能为空',
                validators: {
                    notEmpty: {
                        message: '标题不能为空'
                    },
   					stringLength: {
                        min: 5,
                        max: 30,
                        message: '标题长度5-30个字符'
                    },                                   
                }
            },
			context: {
                message: '二维码内容不能为空',
                validators: {
                    notEmpty: {
                        message: '二维码内容不能为空'
                    },
                    stringLength: {
                        min: 10,
                        max: 50,
                        message: '二维码内容长度10-50个字符'
                    },
                    
                }
            },
			
			height: {
                message: '高度不能为空',
                validators: {
                    notEmpty: {
                        message: '高度不能为空'
                    },
                    between: {
                        min: 50,
                        max: 500,
                        message: '高度50-500之间'
                    },
                    
                }
            },
			
			length: {
                message: '不能为空',
                validators: {
                    notEmpty: {
                        message: '长度不能为空'
                    },
                    between: {
                        min: 50,
                        max: 500,
                        message: '长度50-500之间'
                    },
                    
                }
            },
		
        }
    }) 
		

});
