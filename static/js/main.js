function deleteRes(elem) {
	console.log($(elem))
	console.log({file_name: $(elem)[0].previousElementSibling.children[0].innerHTML})
    $.ajax({
        method: "POST",
        url: window.location,
        data: JSON.stringify({file_name: $(elem)[0].previousElementSibling.children[0].innerHTML}),
        dataType: "json"
    })
    $($(elem)[0].parentNode).remove()
}
