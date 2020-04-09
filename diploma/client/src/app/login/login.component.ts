import { Component, OnInit } from '@angular/core';
import { ApiService } from '../api.service';
import { Router } from '@angular/router';
import { normalizeGenFileSuffix } from '@angular/compiler/src/aot/util';

@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.css']
})
export class LoginComponent implements OnInit {
  public email
  public password
  public conpassword
  public username
  public incorrect = false
  public ischecked = true
  public isequal = true
  public message
  public userType
  public  options = [
    { name: "student", value: 1 },
    { name: "master", value: 2 }
  ]
  constructor(
    private apiService: ApiService,
    private router: Router
  ) { }

  public registration = false;
  ngOnInit() {
   
  }

  user(id, username, password) {
    return {
      id: id,
      username: username,
      password: password,
      created_at: new Date(),
      updated_at: new Date(),
      user_info_id : null,
      usertype: this.userType?this.userType:0
    }
  }

  logIn() {
    this.ischecked = this.checkFill(1)
    if (!this.ischecked) return;
    let user = this.user(1, this.username, this.password)
    this.apiService.signIn(user).subscribe(res => {
      if (res.message == "Succes") {
        sessionStorage.setItem('user',res.userType)
        sessionStorage.setItem('userId',res.user_id)
        this.router.navigate(['bar'])
        this.incorrect = false
      } else {
        this.incorrect = true
      }
    });
  }

  SignUp() {
    this.ischecked = this.checkFill(2)
    if (!this.ischecked) return
    this.isequal = this.password == this.conpassword ? true : false
    if (!this.isequal) return
    let user = this.user(1, this.username, this.password)
    this.apiService.createUsers(user).subscribe(res => {
      if (res.message == 'Succes') {
        this.regClose()
      } else if (res.message == 'Username is exist'){
        this.message = res.message
      }
    });
  }

  checkFill(type) {
    let ischecked;
    if (type == 1) {
      ischecked = !this.password || !this.username ? false : true
    } else if (type == 2) {
      ischecked = !this.password || !this.username || !this.conpassword ? false : true
    }
    return ischecked
  }

  regOpen() {
    this.resetFun(1)
  }

  regClose() {
    this.resetFun(2)
  }

  resetFun(type) {
    this.isequal = this.ischecked = true
    this.incorrect = false
    this.message = null
    this.registration = type == 1 ? true : false
  }
}
