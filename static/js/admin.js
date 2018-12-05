

jQuery(document).ready(function() 
{
		
	$('#categories-view-admin_panel .cat-action').on('click', function(){
			
		alert("here");		

		});
	
$("#add-new-cat form").submit(function(event){
	
		

		
	
		event.preventDefault(); 		
		
		var v_active = false;
		
		var v_name = document.getElementById("new-cat-name").value;
		var v_parent = document.getElementById("new-cat-parent").value;
		
		if ($('#new-cat-active').is(':checked')){
			var v_active = true;
		} 
		
		var v_weight = document.getElementById("new-cat-weight").value;
		var v_url = document.getElementById("new-cat-url").value;
		
		
		/*
		alert(v_name + " - " + v_parent + " - " + v_active + " - " + v_weight + " - " + v_url);
		return;
	*/
	
		var formData = {
						name:v_name,												
						parent_id:v_parent,
						wheight:v_weight,
						active:v_active,
						url:v_url
					}; 
		
		$.ajaxSetup({cache: false}); 	

		$.ajax({
		  type: "POST",
		  url: "/admin/categories/add",
		  data: formData,			  
		  success: function(data){	
			if (data == "ok"){
				location.reload();
			} else {
				alert(data);
			}
			//ShowPersonResults(data);
		  }
		});		
	}); 

		
	
});


	
	








