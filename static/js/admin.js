

jQuery(document).ready(function() 
{
	
	/*$('#progressDiv').progressbar({ value: 21 });
	alert("here");*/
	
	$('#AccessTables').DataTable({
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
			}            
        }        		
	});
	
	
	
	$('#ReloadOISUU').on('click', function(){
			
		
		$(this).css('color','red');
		var formData = {id:0}; 
		
		$.ajaxSetup({cache: false}); 
		
		$.ajax({
			type: "POST",
			url: "/admin/databases/update/oisuu",
			data: formData,			  
			success: function(data){	
							alert(data);	
							$("#ReloadOISUU").css('color','#333');							
					 }
			});
	});
	
	
	$('#AccessTables .Update').on('click', function(){
		var id = $(this).attr('table_id');
		$(this).css('color','red');
		var formData = {
						id:id
					}; 
		
		$.ajaxSetup({cache: false}); 
		
		$.ajax({
			type: "POST",
			url: "/admin/databases/update",
			data: formData,			  
			success: function(data){	
							var res = data.split('|');	
							$("#Update"+res[0]).css('color','#333');
							var index = $("#Update"+res[0]).parent().index();
							
							$("#Update"+res[0]).parent().parent().children('td:eq(4)').text(FormatedDate());
					 }
			});
	});
		
	
});

	
	








