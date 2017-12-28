function subForm(tid) {
    var kw = document.getElementById("kw");
    var qFrom = document.getElementById("queryForm");
    switch (tid) {
        case 1:
            qFrom.action = "http://www.google.com/search";
            kw.name = "q";
            kw.setAttribute("name", "q");
            qFrom.submit();
            break;
        case 2:
            qFrom.action = "http://www.baidu.com/baidu";
            kw.name = "wd";
            kw.setAttribute("name", "wd");
            qFrom.submit();
            break;
        case 3:
            qFrom.action = "http://wwww.bing.com/search";
            kw.name = "q";
            kw.setAttribute("name", "q");
            qFrom.submit();
            break;
    }
}

function submitForm(e) {
    var event = window.event ? window.event : e;
    if (event.keyCode == 13) {
        subForm(1);
    }
}

var params = {
    "XOffset": 0, //提示框位置横向偏移量,单位px
    "YOffset": 0, //提示框位置纵向偏移量,单位px
    "fontSize": "14px",		//文字大小
    "fontFamily": "宋体",	//文字字体
    "borderColor": "gray", 	//提示框的边框颜色
    "bgcolorHI": "#03c",		//提示框高亮选择的颜色
    "sugSubmit": false		//在选择提示词条是是否提交表单
};
try{
    BaiduSuggestion.bind("kw", params, function (kwords) {
        document.getElementById("kw").value = kwords;
        subForm(1);
    });
    document.getElementById("kw").focus();
}catch(err){}