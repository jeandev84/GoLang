$(function () {
	$("#content").bind("input change", function () {
		$.post("/gethtml", {md: $("#content").val()}, function (response) {
			$("#md_html").html(response.html)
		});
	});
})

// md - это параметр который мы прислаем по методу POST