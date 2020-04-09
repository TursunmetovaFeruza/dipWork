import { Component, OnInit } from '@angular/core';
import { ApiService } from '../api.service';
import { ActivatedRoute } from '@angular/router';
import {DropdownModule} from 'primeng/dropdown';
import {SelectItem} from 'primeng/api';
@Component({
  selector: 'app-users',
  templateUrl: './users.component.html',
  styleUrls: ['./users.component.css']
})
export class UsersComponent implements OnInit {
  public items
  public sysNames
  selectedCar: string;

  constructor(
    private apiService: ApiService,
    public activatedRoute: ActivatedRoute
  ) {
    this.activatedRoute.params.subscribe(async ({ type, child }) => {
      this.getSysName(child)
      if (child == 'students') {
        this.apiService.GetAllStudents().subscribe(res => {
          this.items = res.user;
        });
      } else if (child == 'masters') {
        this.apiService.GetAllMaters().subscribe(res => {
          this.items = res.user;
        });
      } else if (child == 'credentials') {
        this.apiService.GetAllUsers().subscribe(res => {
          this.items = res.user;
          this.items.forEach(el => {
            el.usertype = el.usertype && el.usertype == 1?'student':el.usertype && el.usertype == 2?'master':'admin';
            // el.usertype = el.usertype === 0?'admin':null;
          });
        });
      }
    })
   }

 
  ngOnInit(): void {
  // console.log(this.items)
  }

 getSysName(child){
  if (!child) return
  if (child == 'students') {
    this.sysNames = [
      {name:'name', label:'Имя', selected:''},
      {name:'surname', label:'Фамилия', selected:''},
      {name:'fathername', label:'Отчество', selected:''},
      {name:'email', label:'Почта', selected:''},
      {name:'faculty', label:'Факультет', selected:''},
      {name:'school', label:'Кафедра', selected:''},
      {name:'speciality', label:'Специальность', selected:''},
      {name:'group', label:'Группа', selected:''},
      {name:'enrollmentyear', label:'Год поступления', selected:''},
      {name:'lang', label:'Язык обучения', selected:''},
    ]
  } else if (child == 'masters') {
    this.sysNames = [
      {name:'name', label:'Имя', selected:''},
      {name:'surname', label:'Фамилия', selected:''},
      {name:'fathername', label:'Отчество', selected:''},
      {name:'email', label:'Почта', selected:''},
      {name:'faculty', label:'Факультет', selected:''},
      {name:'school', label:'Кафедра', selected:''}
    ]
  } else if (child == 'credentials') {
    this.sysNames = [
      {name:'username', label:'Логин', selected:''},
      {name:'usertype', label:'Тип пользователя', selected:''},
      {name:'created_at', label:'Дата создания', selected:'',type:'date'},
      {name:'updated_at', label:'Дата обновения', selected:'',type:'date'},
      {name:'password', label:'Пароль', selected:''}
    ]
  }
 }

}
