

    <div class="container-fluid">
		<!-- Breadcrumbs-->
		<ol class="breadcrumb">
			<li class="breadcrumb-item"><a href="#">Система</a></li>
			<li class="breadcrumb-item active">Категории</li>
		</ol>
  
		<table id="AccessTables" style="width:100%; font-size: 0.75rem;" class="table table-bordered compact" cellspacing="0">
			<!-- AccessTables -->
			<thead>
				<tr>			   	            
					<th>№</th>					
					<th>Название</th>					
				</tr>				
			</thead>	
			<tbody>
			{{if .Categories -}}						
				{{ range $key, $value := .Categories }}										
				<tr> 				
					<td>{{$key}}</td>
					<td>{{$value.Name}}</td> 				
					<!-- <td><span class="Update" id="Update{{$value.Id}}" table_id="{{$value.Id}}">Обновить</span></td>  -->
				</tr>
			{{end}}						
			{{- else}}
					Не удалось загрузить список категорий
			{{- end}}
			</tbody>
		</table>
	</div>

