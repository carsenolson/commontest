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

function startTesting(elem) {
	console.log($(elem))
	$(elem).html("Stop testing")	
	$(elem).attr("onclick", "stopTesting(this)")	
	$.ajax({
		method: "POST",
		url: "/",
		error: function (jqxhr, stat, error){
			let info = document.createElement("p")
			$(info).html("error occured: "+error)		
			$(info).css("margin-top", "15px")	
			$(elem.parentNode).append(info)	
		},
		data: JSON.stringify({action: "startTesting", option: ""}),
		dataType: "text",
		success: function(data, stat, jqxhr) {
			console.log(data)
			let info = document.createElement("p")	
			$(info).html("Share this URL: "+data)	
			$(info).css("margin-top", "15px")	
			$(elem.parentNode).append(info)	
		}
	})
	
}

function stopTesting(elem) {
	console.log($(elem))
	$(elem).html("Start testing")
	$(elem).attr("onclick", "startTesting(this)")
	$(elem.nextElementSibling).remove()	
	$.ajax({
		method: "POST",
		url: "/",
		data: JSON.stringify({action: "stopTesting", option: ""}),
	})
}
