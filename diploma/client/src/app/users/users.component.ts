import { Component, OnInit } from '@angular/core';
import { ApiService } from '../api.service';
import { ActivatedRoute } from '@angular/router';
@Component({
  selector: 'app-users',
  templateUrl: './users.component.html',
  styleUrls: ['./users.component.css']
})
export class UsersComponent implements OnInit {
  public items
  public sysNames
  public attendance = false
  selectedCar: string;
  public stdate = new Date(2020, 1, 20)
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
            el.usertype = el.usertype && el.usertype == 1 ? 'student' : el.usertype && el.usertype == 2 ? 'master' : 'admin';
            // el.usertype = el.usertype === 0?'admin':null;
          });
        });
      } else if (child == 'events') {
        let userType = sessionStorage.getItem('user')
        let userId = sessionStorage.getItem('userinfoid')
        if (userType == 'admin') {
          this.apiService.GetAllEvents().subscribe(res => {
            this.items = res.user;
          });
        } else if (userType == 'master') {
          this.attendance = true;
          this.apiService.GetAllEventsForMaster(userId).subscribe(res => {
            this.items = res.user;
            this.items.forEach(el => {
              el.upload = this.checktime(el.starttime, el.endtime)
            });
          });
        } else if (userType == 'student') {
          this.apiService.GetAllEventsForStudent(userId).subscribe(res => {
            this.items = res.user;
            this.items.forEach(el => {
              el.attend = el.attend && el.attend.length?el.attend.filter(x=>x == el.studentlist[0]):null
              el.apsent = el.apsent && el.apsent.length?el.apsent.filter(x=>x == el.studentlist[0]):null

            })
          })
        }

      }
    })
  }

  checktime(starttime, endtime) {
    if (new Date(starttime).getHours() == new Date().getHours() || new Date(endtime) >= new Date()) {
      return false
      // (new Date(endtime).getHours()>=new Date().getHours() && new Date(endtime).getMinutes()>=new Date().getMinutes() )
    } else {
      return true
    }
  }

  ngOnInit(): void {
    // console.log(this.items)
  }

  getSysName(child) {
    if (!child) return
    if (child == 'students') {
      this.sysNames = [
        { name: 'name', label: 'Имя', selected: '' },
        { name: 'surname', label: 'Фамилия', selected: '' },
        { name: 'fathername', label: 'Отчество', selected: '' },
        { name: 'email', label: 'Почта', selected: '' },
        { name: 'faculty', label: 'Факультет', selected: '' },
        { name: 'school', label: 'Кафедра', selected: '' },
        { name: 'speciality', label: 'Специальность', selected: '' },
        { name: 'group', label: 'Группа', selected: '' },
        { name: 'enrollmentyear', label: 'Год поступления', selected: '' },
        { name: 'lang', label: 'Язык обучения', selected: '' },
      ]
    } else if (child == 'masters') {
      this.sysNames = [
        { name: 'name', label: 'Имя', selected: '' },
        { name: 'surname', label: 'Фамилия', selected: '' },
        { name: 'fathername', label: 'Отчество', selected: '' },
        { name: 'email', label: 'Почта', selected: '' },
        { name: 'faculty', label: 'Факультет', selected: '' },
        { name: 'school', label: 'Кафедра', selected: '' }
      ]
    } else if (child == 'credentials') {
      this.sysNames = [
        { name: 'username', label: 'Логин', selected: '' },
        { name: 'usertype', label: 'Тип пользователя', selected: '' },
        { name: 'created_at', label: 'Дата создания', selected: '', type: 'date' },
        { name: 'updated_at', label: 'Дата обновения', selected: '', type: 'date' },
        { name: 'password', label: 'Пароль', selected: '' }
      ]
    } else if (child == 'events') {
      this.sysNames = [
        { name: 'id', label: 'ID', selected: '' },
        { name: 'lection', label: 'Лекция', selected: '' },
        { name: 'master', label: 'Преподаватель', selected: '' },
        { name: 'starttime', label: 'Начало', selected: '', type: 'date' },
        { name: 'endtime', label: 'Конец', selected: '', type: 'date' },
        { name: 'studentlist', label: 'Студенты', selected: '', type: 'array' },
        { name: 'attend', label: 'Присутисвовал', selected: '', type: 'array' },
        { name: 'apsent', label: 'Отсутвовал', selected: '', type: 'array' }
      ]
    }
  }

  handleFileInput(id,students,$event) {
    // console.log(stud, $event.target.files[0].type)

    this.apiService.UploadAttendance(id,students, $event.target.files[0]).subscribe(res => {
      console.log(res)
    });
  }
}
