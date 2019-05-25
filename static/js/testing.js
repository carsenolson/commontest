let timer = setTimeout(walkThrough, $(".badge").html() * (1000 * 60))

function scaleImage(elem) {
	$(elem).attr("class", "center-fit")
	$(elem).attr("onclick", "unscaleImage(this)")    
}

function unscaleImage(elem) {
	$(elem).attr("class", "classic-image")
	$(elem).attr("onclick", "scaleImage(this)")
}

function checkAnswer(elem) {
	// Switch colors on click	
	if ($(elem).css("background-color") == "rgb(249, 250, 251)") {
		$(elem).css("background-color", "rgb(40, 167, 69)")
	} else if ($(elem).css("background-color") == "rgb(40, 167, 69)") {
		$(elem).css("background-color", "rgb(249, 250, 251)")
	}
}

function walkThrough(elem) {
	clearTimeout(timer)	
	let data = {
		"Full_name": $("input")[0].value,
		"File_name": $("input")[2].value,
		"Group": $("input")[1].value,
		"Picked_answers": []
	}
	let questions = $(document.body).find(".question"); 
	console.log($("input"))	
			
	// loop over questions
	for (let i = 0; i < questions.length; i++) {
		let answers = $(questions[i]).find(".answer")	
		let true_answers = []
		for (let j = 0; j < answers.length; j++) {
			if ($(answers[j]).css("background-color") == "rgb(40, 167, 69)") {
				true_answers.push(j)	
			}
			continue
		}
		data["Picked_answers"].push(true_answers)
	}
	console.log(JSON.stringify(data))
	
	$.ajax({
		method:"POST",
		url: "/result",
		data: JSON.stringify(data),
		dataType: "text",	
		error: function(jqxhr, stat, error) {
			console.log(error)	
		},
		success: function (data, stat, jqxhr) {
			$("body").html("")	
			$("body").append(`
						<div id="center">
							<h1>`+data+`%</h1>
							<a class="btn btn-primary" href="/" role="button">Try again</a>
						</div>
					`)
			console.log("okay")	
		}
	})
}
