



    <div class="container-fluid">
		<!-- Breadcrumbs-->
		<ol class="breadcrumb">
			<li class="breadcrumb-item"><a href="#">Система</a></li>
			<li class="breadcrumb-item active">Категории</li>
		</ol>
  
		<div class="row">
			<div class="col-md-7">
				
				{{.CategoriesRender}}
				
			</div>
			<div class="col-md-5">
				<ul class="nav nav-tabs" id="CatTab" role="tablist">
				  <li class="nav-item">
					<a class="nav-link active" href="#add-new-cat" data-toggle="tab">Новая</a>
				  </li>
				  <li class="nav-item">
					<a class="nav-link" href="#edit-cat" data-toggle="tab">Редактировать</a>
				  </li>				  
				</ul>
				
				<div class="tab-content" style="padding-top:25px;">	
				
					<div class="tab-pane active" id="add-new-cat">
						<h4>Добавить новую</h4>
						<form >				
							<div class="form-group row">
								<label for="cat-name" class="col-sm-2 col-form-label" >Название:</label>
								<div class="col-sm-10">  					
								  <input type="text" class="form-control" id="cat-name">
								</div>		
							</div>

							<div class="form-group row">
								<label for="cat-name" class="col-sm-2 col-form-label">Родитель:</label>
								<div class="col-sm-10">
									{{.CategoriesSelectRender}}
								</div>
							</div>
							
							<div class="form-group row">
								<label for="cat-active" class="col-sm-2 col-form-label">Активный:</label>
								<div class="col-sm-10">
									<input type="checkbox" id="cat-active" class="cat-checkbox">
								</div>
							</div>

							<div class="form-group row">
								<label for="cat-weight" class="col-sm-2 col-form-label">Вес:</label>
								<div class="col-sm-10">
									{{.CategoriesSelectWeight}}
								</div>
							</div>
							  
							<div class="form-group row">
								<label for="cat-url" class="col-sm-2 col-form-label">URL:</label>
								<div class="col-sm-10">
									<input type="text" class="form-control" id="cat-url">
								</div>
							</div>
						
							<div class="form-group row">
								<div class="col-sm-12">
								  <button type="submit" 
										class="btn btn-primary float-sm-right"
										onClick="return ValidPerson(this.form);" >Добавить</button>
								</div>
							  </div>
							
						</form>
					</div>
					
					<div class="tab-pane" id="edit-cat">
						<h4>Редактировать</h4>
					</div>
				
				</div>
				
				
			</div>
		</div>
	
	</div>

