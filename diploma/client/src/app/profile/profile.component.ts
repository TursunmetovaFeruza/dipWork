import { Component, OnInit, Input } from '@angular/core';
import { ApiService } from '../api.service';
import { DropdownModule } from 'primeng/dropdown';

interface User {
  name: string;
  surname: string;
  fathername: string;
  enrollmentyear: number;
  courseId: number;
  langId: number;
  groupId: number;
  email: string;
  lectionId: number;
  facultyId: number;
  schoolId: number;
  specialityId: number
}


@Component({
  selector: 'app-profile',
  templateUrl: './profile.component.html',
  styleUrls: ['./profile.component.css']
})
export class ProfileComponent implements OnInit {

  constructor(
    private apiService: ApiService,

  ) { }
  public faculty
  public school
  public name
  public surname
  public fathername
  public speciality
  public enyear
  public email
  public username
  public password
  public group
  public course
  public lang
  public user
  public isdisable = true
  public isStudent = false
  public isMaster = false
  public inputype
  public isCreated
  public newUsers: User
  @Input() addtype: string

  @Input() _date: any
  @Input()
  set date(date: any) {
    if (date) {
      this.isCreated = true;
      this.isMaster = this.addtype == 'masters' ? true : false
      this.isStudent = this.addtype == 'students' ? true : false
      this.user = null;
      this.setPermitions()
      this.getTables()
    }
  }
  // public 
  ngOnInit(): void {
    this.setPermitions()
    this.getUser()
  }

  getUser() {
    let userType = sessionStorage.getItem('user')
    let id = sessionStorage.getItem('userId')
    this.apiService.getUser(id, userType).subscribe(res => {
      this.user = res.user
    });

  }

  setPermitions() {
    let userType = sessionStorage.getItem('user')
    if (userType == 'admin') {
      this.isdisable = false
    } else if (userType == 'student') {
      this.isStudent = true
    } else if (userType == 'master') {
      this.isMaster = true
    }
  }
  newItem() {
    let User = {
      name: this.name,
      surname: this.surname,
      fathername: this.fathername,
      enrollmentyear: parseInt(this.enyear),
      courseId: parseInt(this.course),
      langId: parseInt(this.lang),
      groupId: parseInt(this.group),
      email: this.email,
      lectionId: null,
      facultyId: parseInt(this.faculty),
      schoolId: parseInt(this.school),
      specialityId: parseInt(this.speciality)
    }
    let userform = {
      id: 0,
      username: this.username,
      password: this.password,
      created_at: new Date(),
      updated_at: new Date(),
      user_info_id: null,
      usertype: this.addtype == 'students' ? 1 : 2
    }

    if (this.addtype == 'students') {
      this.createStudent(User, userform)
    } else if (this.addtype == 'masters') {
      this.createMaster(User, userform)
    } else {
      this.createCredentials(userform)
    }
    console.log(User)
  }

  createStudent(User, userform) {
    this.apiService.CreateStudent(User).subscribe(res => {
      if (res.message == 'Success') {
        userform.user_info_id = parseInt(res.id)
        this.createCredentials(userform)
      }
    })
  }

  createMaster(User,userform) {
    this.apiService.CreateMaster(User).subscribe(res => {
      if (res.message == 'Success') {
        userform.user_info_id = parseInt(res.id)
        this.createCredentials(userform)
      }
    })
  }

  createCredentials(userform) {
    this.apiService.createUsers(userform).subscribe(res => {
      if (res.message == 'Succes') {
        console.log('good')
      } else if (res.message == 'Username is exist') {
        console.log('bad')
      }
    })
  }

  public faculties
  public schools
  public specialities
  public courses
  public groups
  public langs
  // public lections
  public usertypes
  getTables() {
    this.apiService.GetFaculties().subscribe(res => {
      this.faculties = res.user
    });
    this.apiService.GetSchools().subscribe(res => {
      this.schools = res.user
    });
    this.apiService.GetSpecialities().subscribe(res => {
      this.specialities = res.user
    });
    this.apiService.GetCourses().subscribe(res => {
      this.courses = res.user
    });
    this.apiService.GetGroups().subscribe(res => {
      this.groups = res.user
    });
    this.apiService.GetLanguages().subscribe(res => {
      this.langs = res.user
    });
    // this.apiService.GetLections().subscribe(res => {
    //   this.lections = res.user
    // });
    this.apiService.GetUserType().subscribe(res => {
      this.usertypes = res.user
    });

  }
  handleFileInput($event){
    console.log($event.target.files[0].type)
    
    this.apiService.UploadFile($event.target.files[0]).subscribe(res => {
      console.log(res)
    });
  }
}
