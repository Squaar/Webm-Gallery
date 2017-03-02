var serviceRoot = location.origin;
var files;
var fileGrid;

$(document).ready(function(){
	$.getJSON(serviceRoot + '/files', function(data){
		files = data;
		buildGrid();
	})
	.fail(function(){
		console.log('Failed to get available files');
	});
})

// function buildGrid(){
// 	var nCols = 6;
// 	var nRows = files.length / nCols;
// 	var gridHTML = "";

// 	fileGrid = [];
// 	var i = 0;
// 	for(var r=0; r<nRows; r++){
// 		gridHTML += '<div class="row">'
// 		fileGrid[r] = [];
// 		for(var c=0; c<nCols && i<files.length; c++){
// 			//gridHTML += '<div class="embed-responsive embed-responsive-16by9 col-md-' + 12/nCols + '">' + 
// 				//'<video class="embed-responsive-item" controls><source src="' + serviceRoot + '/file/' + files[i] + '"></video>'
// 			gridHTML += '<div class="col-md-' + 12/nCols + '">' + 
// 				'<h5><a href="' + serviceRoot + '/file/' + files[i] + '">' + files[i] + '</a></h5></div>';

// 			fileGrid[r][c] = files[i];
// 			i++;
// 		}
// 		gridHTML += '</div>';
// 	}

// 	$('#grid').append(gridHTML);
// }

function buildGrid(){
	var nCols = 6;
	var gridHTML = '<div class="row no-gutters">';

	var i = 0;
	for(var i=0; i<files.length; i++){
		//gridHTML += '<div class="embed-responsive embed-responsive-16by9 col-md-' + 12/nCols + '">' + 
			//'<video class="embed-responsive-item" controls><source src="' + serviceRoot + '/file/' + files[i] + '"></video>'
		gridHTML += '<div class="cell col-xs-6 col-md-' + 12/nCols + '">' + 
			'<div><h5><a href="' + serviceRoot + '/file/' + files[i] + '"><img src="/img/SpicyDancer.png"></a></h5></div>' +
			'</div>';
	}
	gridHTML += '</div>';
	$('#grid').append(gridHTML);
}