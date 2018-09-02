$(function () {
    var url = $(location).attr('pathname');
    var paths = url.split("/");
    if (paths.length >= 1 && $("#nav_item_" + paths[1]).length) {
        $("#nav_item_" + paths[1]).addClass("active");
    } else {
        $("#nav_item_index").addClass("active");
    }
    //custom copyrightText
    $("#footerCopyrightText").html("&copy; BLOG.IOSXC.COM &middot; "+new Date().getFullYear());
});