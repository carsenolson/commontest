questionIndex = $(".accardion").children().length+1;
console.log(questionIndex);

function handleFileSelect(evt) {
    let reader = new FileReader()
    reader.onload = (function(theFile) {
        return function(e) {
            let img = document.createElement("img")
            $(img).attr('src', e.target.result)
            $(img).attr('title', theFile.name)
            $(img).css("height", "70px")
            $($(evt)[0].nextElementSibling.nextElementSibling).html("")
            $($(evt)[0].nextElementSibling.nextElementSibling).append(img)
        }
        ;
    }
    )(evt.files[0]);
    reader.readAsDataURL(evt.files[0])
}

$("#save-test").on("click", walkThrough);

function walkThrough() {
    // Create test data structure then parse everything	
    let data = {
        File_name: "",
        Test: {
            Name: "",
            Time: 20,
            Questions: []
        }
    }
    readers = [];
    let test_option_inputs = $(".input-group-custom").find("input")
    data["File_name"] = test_option_inputs[0].value
    data["Test"]["Name"] = test_option_inputs[1].value
    data["Test"]["Time"] = test_option_inputs[2].value
    // Walk through questions
    for (question of $(".accordion").find(".question")) {
        let qt = {
            Title: "",
            Answers: [],
            Image: [],
            True_answers: []
        }
        qt["Title"] = question.firstElementChild.value
        //Title		
        // loop over answers	
        let answers = $(question).find(".answer-area").find("input");
        for (let i = 0; i < answers.length; i++) {
            if ($(answers[i]).attr("class") == "right-answer" && answers[i].checked == true) {
                if (i == 0) {
                    qt["True_answers"].push(i)
                } else if (i == 2) {
                    qt["True_answers"].push(i - 1)
                } else {
                    qt["True_answers"].push(i / 2)
                }
            }
            if ($(answers[i]).attr("class") == "form-control") {
                qt["Answers"].push(answers[i].value)
            }
        }
        // loop over images
        for (image of $(question).find(".image-area").find("img")) {
            qt["Image"].push([$(image).attr('title'), $(image).attr('src')])
        }
        data["Test"]["Questions"].push(qt)
    }
    sendPostData(data)
}

function deleteQuestion(elem) {
    $($(elem))[0].parentNode.parentNode.parentNode.remove()
}

function addImage(elem) {
    $($(elem)[0].parentNode.previousElementSibling.previousElementSibling).append(`
						<div class="input-group d-inline">	
							<input type="file" class="form-control-file-sm" style="width:80%" onchange="handleFileSelect(this)">
                            	<button class="btn btn-sm btn-outline-danger delete-answer" type="button" id="button-addon2" 
																onclick="deleteImage(this)" style="width: 17%">Del</button>
							<output class="list"></output><br>	
						</div> 
			`)
}

function deleteImage(elem) {
    $($(elem))[0].parentNode.remove()
}

function rightAnswer(elem) {
    if ($($(elem))[0].checked)
        $($(elem)[0].parentNode).css("background", "#28a745");
    else if (!$($(elem))[0].checked)
        $($(elem)[0].parentNode).css("background", "#e9ecef");
}

function sendPostData(data) {
	$.ajax({
        method: "POST",
        url: "/newtest",
        data: JSON.stringify(data),
       	error: function(jqxhr, textStatus, error) {
			console.log(textStatus);
			console.log(error);
		},
		dataType: "text",
   		success: function(data, status, jqxhr) {
			if (data == "okay")	
			location.href = "/"	
		}
	});
}

function deleteAnswer(elem) {
    $($(elem))[0].parentNode.parentNode.remove()
}

function addAnswer(elem) {
    $($(elem)[0].parentNode.previousElementSibling).append(`
                        <div class="answer">
						 							<div class="input-group">
														<div class="input-group-text">
															<input type="checkbox" class="right-answer" onclick="rightAnswer(this)" aria-label="Checkbox for following text input">
														</div> 
														<input type="text" class="form-control" placeholder="Recipient's username" 
															aria-label="Possible answer" aria-describedby="button-addon2">
														<div class="input-group-append">
															<button class="btn btn-outline-danger delete-answer" type="button" 
																	onclick="deleteAnswer(this)"id="button-addon2">Del</button>
														</div>
													</div>
												</div>
														`);
}

// Handler for add question button
$('#add-question').on("click", function(event) {
    questionIndex++;
    $(".accordion").append(`
		<div class="card">
      <div class="card-header" id="Question">
        <h2 class="mb-0">
            <button class="btn btn-link" type="button" data-toggle="collapse" data-target="#collapse` + questionIndex + `" aria-expanded="true" aria-controls="collapseOne"> Question  
                  </button>
            <button type="button" class="btn btn-danger float-right delete-question" onclick="deleteQuestion(this)">Delete</button>
        </h2>
    </div>
    <div id="collapse` + questionIndex + `" class="collapse show" aria-labelledby="headingOne" data-parent="#accordionExample">
        <div class="card-body">
            <div class="question">
                <input type="text" class="form-control" placeholder="Question title">
                <div class="image-area" style="margin-top: 20px;">
                    <div class="input-group d-inline">
                        <input type="file" class="form-control-file-sm" style="width: 80%" onchange="handleFileSelect(this)">
                        <button class="btn btn-sm btn-outline-danger delete-answer" type="button" id="button-addon2" onclick="deleteImage(this)" style="width: 17%">Del</button>
                        <output class="list"></output><br>
                    </div>
                </div>
                <div class="answer-area" style="margin: 20px 0px">
                    <div class="answer">
                        <div class="input-group">
                            <div class="input-group-text">
                                <input type="checkbox" class="right-answer" onclick="rightAnswer(this)" aria-label="Checkbox for following text input">
                            </div>
                            <input type="text" class="form-control" placeholder="Possible answer" aria-label="Recipient's username" aria-describedby="button-addon2">
                            <div class="input-group-append">
                                <button class="btn btn-outline-danger delete-answer" type="button" id="button-addon2" onclick="deleteAnswer(this)">Del</button>
                            </div>
                        </div>
                    </div>
                </div>
                <div class="btn-group" role="group">
                    <button type="button" class="btn btn-sm btn-outline-secondary" onclick="addImage(this)">Add image</button>
                    <button type="button" class="btn btn-sm btn-outline-secondary" onclick="addAnswer(this)">Add answer</button>
                </d-inlinezv>
            </div>
        </div>						
      `);
});
