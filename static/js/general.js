function deleteTest(elem) {
	console.log($(elem))
	$.ajax({
		method: "POST",
		url: "/",
		data: JSON.stringify({action: "deleteTest", option: $($(elem)[0].parentNode.parentNode).find("h5").html()}), 
		dataType: "json"
	})
	$($(elem)[0].parentNode.parentNode).remove()
}
