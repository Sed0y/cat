
// https://www.avito.ru/user/19463835664bb87bb5f4642980d2840b/profile?id=1391194144&src=item

function IhDateToNashDate(ih_date){
	
	//alert();
	var ih_date = res.split('-');		
		
	var year = ih_date[0];
	var month = ih_date[1];		
	var day = ih_date[2];	
			
	return day + "." + month + "." + year;
};

function ToMoney(str, currency) {
	
	if (currency == "" || currency == undefined)
		currency = "р.";
    var parts = (str + '').split('.'),
        main = parts[0],
        len = main.length,
        output = '',
        i = len - 1;
    
    while(i >= 0) {
        output = main.charAt(i) + output;
        if ((len - i) % 3 === 0 && i > 0) {
            output = ' ' + output;
        }
        --i;
    }

    if (parts.length > 1) {
        output += '.' + parts[1];
    }
    return output + "" + currency;
};

function ToMoney2(str) {
	
    var parts = (str + '').split('.'),
        main = parts[0],
        len = main.length,
        output = '',
        i = len - 1;
    
    while(i >= 0) {
        output = main.charAt(i) + output;
        if ((len - i) % 3 === 0 && i > 0) {
            output = ',' + output;
        }
        --i;
    }

    if (parts.length > 1) {
        output += ',' + parts[1];
    }
    return output;
};

function FormatedDate() {	
	
	
	Data = new Date();
	
	Year = Data.getFullYear()+"";
	
	Month = (Data.getMonth() + 1)+"";
	if( Month.length == 1){
		Month = "0" + Month;
	}
	
	Day = Data.getDate() + "";
	if( Day.length == 1){
		Day = "0" + Day;
	}
	
	Hour = Data.getHours() +"";
	Minutes = Data.getMinutes() +"";
	
	if( Minutes.length == 1){
		Minutes = "0" + Minutes;
	}
	
    return Day + "." + Month + "." +  Year + " " +  Hour + ":" +  Minutes;
	
};

function FormatDateTo_ymd(date) {
	
	var parts = (date + '').split('-');
	
	return parts[2] + '-' + parts[1] + '-' + parts[0]
	
}

function Translate(str) {
	
	
	switch (str) {
	  case "loss":
		return "Убыток";
	  case "zone":
		return "Филиал";
	  case "ins_type":
		return "Вид";
	  case "trass_date":
		return "Заявка на трасологию";
	  case "trass_result_date":
		return "Дата результата";
	  case "trass_result":
		return "Результат по заявке";
	  case "finanse":
		return "Сумма минимизации";
	  case "employer_trass":
		return "Работал";
	  case "employer_result":
		return "Куратор";
	  case "date":
		return "Дата";
	  case "filial":
		return "Филиал";
	  case "insurer":
		return "Страхователь";
	  case "vin":
		return "VIN номер";		
	  case "result":
		return "Результат";	
	  case "user":
		return "Сотрудник";	
	  case "vehicle":
		return "Модель";	
	  case "created":
		return "Создан";	
	  case "closed":
		return "Закрыт";	
	  case "usherb":
		return "Ущерб";
	  case "money":
		return "Фин. рез.";
	  case "control":
		return "Контроль";
		default:
		return str;
	}
	
    //return str;
};