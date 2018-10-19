<!DOCTYPE html>
<html lang="en">

<head>
  <meta charset="utf-8">
  <meta http-equiv="X-UA-Compatible" content="IE=edge">
  <meta name="viewport" content="width=device-width, initial-scale=1, shrink-to-fit=no">
  <meta name="description" content="">
  <meta name="author" content="">
  <title>Панель управления</title>
  <!-- Bootstrap core CSS-->
  <link href="/static/js/vendor/bootstrap/css/bootstrap.min.css" rel="stylesheet">
  <!-- Custom fonts for this template-->
  <link href="/static/js/vendor/font-awesome/css/font-awesome.min.css" rel="stylesheet" type="text/css">
  <!-- Page level plugin CSS-->
  <link href="/static/js/vendor/datatables/dataTables.bootstrap4.css" rel="stylesheet">
  <!-- Custom styles for this template-->
  <link href="/static/css/sb-admin.css" rel="stylesheet">      
  <link href="/static/css/admin.css" rel="stylesheet">

  <link href="/static/css/jquery-ui.theme.min.css" rel="stylesheet">
  <link href="/static/css/jquery-ui.structure.min.css" rel="stylesheet">
		
		
</head>

<body class="fixed-nav sticky-footer bg-dark" id="page-top">

	{{.Navigation}}	
 
  <div class="content-wrapper">
	
	{{.LayoutContent}}

	
    <!-- /.container-fluid-->
    <!-- /.content-wrapper-->
    <footer class="sticky-footer">
      <div class="container">
        <div class="text-center">
          <small>Cat v.0.0.1</small>
        </div>
      </div>
    </footer>
    <!-- Scroll to Top Button-->
    <a class="scroll-to-top rounded" href="#page-top">
      <i class="fa fa-angle-up"></i>
    </a>
    <!-- Logout Modal-->
    <div class="modal fade" id="exampleModal" tabindex="-1" role="dialog" aria-labelledby="exampleModalLabel" aria-hidden="true">
      <div class="modal-dialog" role="document">
        <div class="modal-content">
          <div class="modal-header">
            <h5 class="modal-title" id="exampleModalLabel">Уверены, что хотите выйти?</h5>
            <button class="close" type="button" data-dismiss="modal" aria-label="Close">
              <span aria-hidden="true">×</span>
            </button>
          </div>
          <div class="modal-body">Для выхода из системы, нажмите "Выход"</div>
          <div class="modal-footer">
            <button class="btn btn-secondary" type="button" data-dismiss="modal">Отмена</button>
            <a class="btn btn-primary" href="/logout">Выход</a>
          </div>
        </div>
      </div>
    </div>
    <!-- Bootstrap core JavaScript-->
    <script src="/static/js/vendor/jquery/jquery.min.js"></script>
    <script src="/static/js/vendor/bootstrap/js/bootstrap.bundle.min.js"></script>
    <!-- Core plugin JavaScript-->
    <script src="/static/js/vendor/jquery-easing/jquery.easing.min.js"></script>
	<script src="/static/js/extension/jquery-ui.min.js"></script>	
    <!-- Page level plugin JavaScript-->    
    <script src="/static/js/vendor/datatables/jquery.dataTables.js"></script>
    <script src="/static/js/vendor/datatables/dataTables.bootstrap4.js"></script>
    <!-- Custom scripts for all pages-->
    <script src="/static/js/sb/sb-admin.js"></script>	
	<script src="/static/js/admin.js"></script>
	
    <!-- Custom scripts for this page-->
    <!-- <script src="/static/js/sb/sb-admin-datatables.js"></script> -->
    
	<!-- 0aj8q6X9 ></script>		 -->	
  
  
  </div>
</body>

</html>
