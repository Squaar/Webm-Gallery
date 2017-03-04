var files;
var currentFile;

$(document).ready(function(){
	$.getJSON('/files', function(data){
		files = data;
		buildGrid();
	}).fail(function(){
		console.log('Failed to get available files');
	});

	$('#slideshowModal').on('hide.bs.modal', function(e){
		$('.modal-body').empty();
	});
})

function buildGrid(){
	var nCols = 6;
	var gridHTML = '<div class="row">';

	for(var i=0; i<files.length; i++){
		gridHTML += '<div class="cell col-6 col-md-' + 12/nCols + '">' + 
			'<a class="thumbnail" href="javascript:showModal(' + i + ');"><img src="/thumb/' + files[i] + '"></a>' +
			'</div>';
	}
	gridHTML += '</div>';
	$('#grid').append(gridHTML);
}

function showModal(i){
	currentFile = i;
	var videoHTML = '<video class="embed-responsive-item" controls><source src="/file/' + files[i] + '"></video>';
	$('.modal-title').text(files[i]);
	$('.modal-body').append(videoHTML);
	$('#slideshowModal').modal('show');
}
