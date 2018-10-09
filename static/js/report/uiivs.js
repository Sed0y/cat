

jQuery(document).ready(function() 
{
	//'yyyy-mm-dd'
	$('#begin-data-input').dcalendarpicker({theme: 'green', format: 'dd-mm-yyyy' });
	$('#end-data-input').dcalendarpicker({theme: 'green', format: 'dd-mm-yyyy' });
		
	$('#uiivs-report #wrap_body table tr td.click').on('click', function(){		
		
		//alert("click"); 		return;
		
		var attributes = $(this).attr('id').split('-');		
		
		var p_instype = attributes[0];
		var p_restype = attributes[1];
		
		var	p_filial_id = $('#filials').val();

		var formData = {
						id:p_filial_id,												
						begin:FormatDateTo_ymd($('#begin-data-input').val()),
						end:FormatDateTo_ymd($('#end-data-input').val()),						
						instype:p_instype,
						restype:p_restype,
						curator: $('#curator').prop('checked'),
					}; 
						
		$.ajaxSetup({cache: false}); 
		
		$.ajax({
		  type: "POST",
		  url: "/data/uiivs/reguliration/details/filial",
		  data: formData,			  
		  success: function(data){
			//alert(data);
			PrintDetailData(data);
		  }
		});		
	});
	
	
	$('#calculate_uiiivs').on('click', function() {
			
		//alert("click");
		$("#UpmaDetails").text("");
		//ClearReport();									
		//$("#console").append("<p>" + employer_id + "</p></br>");
		var formData = {						
					id: $('#filials').val(),
					curator: $('#curator').prop('checked'),
					begin: FormatDateTo_ymd($('#begin-data-input').val()),
					end: FormatDateTo_ymd($('#end-data-input').val())
				}; 
		
		$.ajaxSetup({cache: false}); 
		$.ajax({
		  type: "POST",
		  url: "/data/uiivs/reguliration/filial",
		  data: formData,			  
		  success: function(data){	
			//alert(data);
			PrintUiivsReportData(data);
		  }
		});
			
	});
	
	
});


function PrintUiivsReportData(data){
	//alert("here2");
	//
	var Results = $.parseJSON(data);
		
	// InsType - вид страхования
	// FilialResults - результат
	// Indicator - параметр
	
	//for (var JobType in Results) {
		//alert(JobType);
	var all_money = 0
	for (var InsType in Results) {
		//alert( InsType);
		
		//for (var FilialResults in Results[InsType]) {
			//alert( InsType + " - " + FilialResults);
			
			for (var Parameter in Results[InsType]) {
				//alert(Parameter);
				//alert( InsType + " - " + FilialResults + " - " + Parameter + "-" + Results[InsType][FilialResults][Parameter]);	
				if( Parameter != "Financial"){
					//$('#' + JobType + '-' + InsType + '-' + Parameter).text(Results[JobType][InsType][EmployerResults][Parameter]+"")
					//alert(Results[InsType][Parameter]);
					$('#' + InsType + '-' + Parameter).text(Results[InsType][Parameter]+"");
				}
				else{
					//alert(Results[InsType][Parameter]);
					//$('#' + JobType + '-' + InsType + '-' + Parameter).text(ToMoney(Results[JobType][InsType][EmployerResults][Parameter]+""))
					$('#' + InsType + '-' + Parameter).text(ToMoney(Results[InsType][Parameter]+""));
					all_money = all_money + 1*Results[InsType][Parameter];
				}
			}
			
		
		
	}	

	
	$('#All').text("0");
	$('#Refusal').text("0");
	$('#Minimization').text("0");
	
	var table = document.getElementById("uiivs_table");
	
	for (var i = 0, row; row = table.rows[i]; i++) {
		
		if (i < 2 || i > 9) { 
			continue;
		}
		
		for (var j = 0, col; col = row.cells[j]; j++) {
			
			if( j == 1) 				
				$('#All').text( 1*$('#All').text() + 1*row.cells[j].innerHTML );
			if( j == 2) 				
				$('#Refusal').text( 1*$('#Refusal').text() + 1*row.cells[j].innerHTML );	
			if( j == 3) 				
				$('#Minimization').text( 1*$('#Minimization').text() + 1*row.cells[j].innerHTML );
		}
		
		$('#Financial').text( ToMoney(all_money)+"" );
		
	}
	

/*

	$('#All').text(
								1*$('#OSAGO-All').text() + 
								1*$('#KASKO-All').text() + 
								1*$('#GO-All').text() + 
								1*$('#ZK-All').text()
							);
	$('#Refusal').text(
								1*$('#OSAGO-RefusalAndMinimization').text() + 
								1*$('#KASKO-RefusalAndMinimization').text() + 
								1*$('#GO-RefusalAndMinimization').text() + 
								1*$('#ZK-RefusalAndMinimization').text()
							);
	$('#Financial').text(
								ToMoney(all_money)+""
							);	
	*/
}


function PrintDetailData(data){
		
	$("#UpmaDetails").text("");
	
	var Results = $.parseJSON(data);
	var render = "";
	var fin_column = -1;
	var id_column = -1;
	
	render += "<table id=\"DetailsTable\" style=\"width:100%\" class=\"table table-bordered compact\" cellspacing=\"0\">";
	render += "	<thead>";
	render += "		<tr>";
	
	for (var part in Results["header"]) {
		if (Results["header"][part] == "finanse" || Results["header"][part] == "money"){
			render += "			<td type=\"money\">" + Translate(Results["header"][part]) + "</td>";
			fin_column = part*1;
		}
		else if (Results["header"][part] == "id"){		
			id_column = part*1;
		}else{		
			render += "			<td>" + Translate(Results["header"][part]) + "</td>";
		}
	}
	
	render += "  </tr>";
	render += "</thead>";	
	render += "<tbody>";
	
	var index
	for (var part in Results["data"]) {			
		render += "<tr>";
		
		index = 0
		for (var val in Results["data"][part]) {
			
			if(index == id_column){
				index++;
				continue;
			}
			
			if (Results["data"][part][val].indexOf("T00:00:00Z") !== -1) {
				res = Results["data"][part][val].replace("T00:00:00Z","");
				
				render += "<td>" + IhDateToNashDate(res) + "</td>";		
				
			} else if(Results["data"][part][val].indexOf("Зона ф-ла") !== -1) {
				res = Results["data"][part][val].replace("Зона ф-ла","");
				render += "<td>" + res + "</td>";
			} else {
				render += "<td>" + Results["data"][part][val] + "</td>";
			}				
			index++;
		}	
		render += "</tr>";
		
	}
	render += "</tbody>";	
	render += "</table>";
	
	if (id_column != -1){
		fin_column = fin_column -1;
	}
	
	$("#UpmaDetails").append(render);
	
	list= document.getElementById("DetailsTable").getElementsByTagName("tr"); 

	for (var i=1; i< list.length; ++i) 
	{ 
		list[i].getElementsByTagName("td")[fin_column].innerHTML= ToMoney2(list[i].getElementsByTagName("td")[fin_column ].innerHTML);  
	}; 

	
	$('#DetailsTable').DataTable({
		"language": {
			"search": "Поиск:",
            "lengthMenu": "По _MENU_ записей на страницу",			
            "zeroRecords": "Ничего не найдено",
            "info": "Страница _PAGE_ из _PAGES_",
            "infoEmpty": "Нет записей",
			"paginate": {
				"first":      "Первая",
				"previous":   "Предыдущая",
				"next":       "Следующая",
				"last":       "Последняя"
			},            
        }
		
	});
	
	//alert("here");
	
}	








