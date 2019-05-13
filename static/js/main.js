function deleteRes(elem) {
	console.log($(elem))
	console.log({file_name: $(elem)[0].previousElementSibling.previousElementSibling.children[0].innerHTML, option:"delete"})
    $.ajax({
        method: "POST",
        url: window.location,
        data: JSON.stringify({file_name: $(elem)[0].previousElementSibling.children[0].innerHTML}),
        dataType: "json"
    })
    $($(elem)[0].parentNode).remove()
}

function xlsxResult(elem) {
	console.log($(elem))
	console.log({file_name: $(elem)[0].previousElementSibling.children[0].innerHTML, option:"xlsx"})
    $.ajax({
        method: "POST",
        url: window.location,
        data: JSON.stringify({file_name: $(elem)[0].previousElementSibling.children[0].innerHTML, option:"xlsx"}),
        dataType: "json"
    })
}
