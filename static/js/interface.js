questionIndex = $(".accardion").children().length;

function deleteQuestion(elem) {	
	$($(elem))[0].parentNode.parentNode.parentNode.remove()
}

function addImage(elem) {
	console.log($($(elem)))	
	console.log($($(elem)[0].previousElementSibling))	
	$($(elem)[0].parentNode.previousElementSibling.previousElementSibling).append(`
						<div class="input-group d-inline">	
							<input type="file" class="form-control-file-sm" style="width:80%">
                            	<button class="btn btn-sm btn-outline-danger delete-answer" type="button" id="button-addon2" 
																onclick="deleteImage(this)" style="width: 17%">Del</button>
						</div> `)
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

$("#save-test").on("click", function(event) {
	console.log("Pressed", event);
});

function deleteAnswer(elem) {
	$($(elem))[0].parentNode.parentNode.remove()
}

function addAnswer(elem) {
	$($(elem)[0].parentNode.previousElementSibling).append(`<div class="answer">
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
                  <button class="btn btn-link" type="button" data-toggle="collapse" data-target="#collapse`+questionIndex+`" aria-expanded="true" aria-controls="collapseOne"> Question  
                  </button>
                  <button type="button" class="btn btn-danger float-right delete-question" onclick="deleteQuestion(this)">Delete</button>      
                </h2>
              </div>
              <div id="collapse`+questionIndex+`" class="collapse show" aria-labelledby="headingOne" data-parent="#accordionExample">
                <div class="card-body">
                  <div class="question">
                      <input type="text" class="form-control" placeholder="Question title">
                      <div class="image-area" style="margin-top: 20px;"> 
					  	<div class="input-group d-inline">	
							<input type="file" class="form-control-file-sm" style="width: 80%">
                            	<button class="btn btn-sm btn-outline-danger delete-answer" type="button" id="button-addon2" 
																onclick="deleteImage(this)" style="width: 17%">Del</button>
						</div> 	
					  </div> 
					  <div class="answer-area" style="margin: 20px 0px">
                          <div class="answer">
                              <div class="input-group">
                                  <div class="input-group-text">
                                      <input type="checkbox" class="right-answer" onclick="rightAnswer(this)" aria-label="Checkbox for following text input">
                                  </div> 
                                  <input type="text" class="form-control" placeholder="Possible answer" aria-label="Recipient's username" 
                                                                                      aria-describedby="button-addon2">
                                  <div class="input-group-append">
                                      <button class="btn btn-outline-danger delete-answer" type="button" id="button-addon2" onclick="deleteAnswer(this)">Del</button>
                                  </div>
                              </div>
                          </div>
                      </div>
                   	  <div class="btn-group" role="group">   
						  <button type="button" class="btn btn-sm btn-outline-secondary" onclick="addImage(this)">Add image</button>
					      <button type="button" class="btn btn-sm btn-outline-secondary" onclick="addAnswer(this)">Add answer</button>
					</div>
                </div>
              </div>	
						`);	
});
