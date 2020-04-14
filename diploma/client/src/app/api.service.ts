import { Injectable } from '@angular/core';
import { HttpClient, HttpHeaders } from '@angular/common/http';
declare var require: any
let qs = require('qs');
import * as _ from "lodash";
// var proxy = require('http-proxy-middleware');



@Injectable()
export class ApiService {

  constructor(private http: HttpClient, ) {
  }
  GetAllUsers() {
    return this.http.get<any>('api/GetAllUsers');
  }

  GetAllStudents() {
    return this.http.get<any>('api/GetAllStudents');
  }

  GetAllMaters() {
    return this.http.get<any>('api/GetAllMaters');
  }

  GetFaculties() {
    return this.http.get<any>('api/GetFaculties');
  }

  GetSchools() {
    return this.http.get<any>('api/GetSchools');
  }

  GetSpecialities() {
    return this.http.get<any>('api/GetSpecialities');
  }

  GetCourses() {
    return this.http.get<any>('api/GetCourses');
  }

  GetLanguages() {
    return this.http.get<any>('api/GetLanguages');
  }

  GetGroups() {
    return this.http.get<any>('api/GetGroups');
  }

  GetUserType() {
    return this.http.get<any>('api/GetUserType');
  }

  GetLections() {
    return this.http.get<any>('api/GetLections');
  }

  createUsers(value): any {
    return this.http.post<any>('api/users', value);
  }
  CreateStudent(value): any {
    return this.http.post<any>('api/CreateStudent', value);
  }
  CreateMaster(value):any {
    return this.http.post<any>('api/CreateMaster',value)
  }
  signIn(user): any {
    return this.http.post<any>('api/user', user)
  }
  getMenuItems(user): any {
    let query = qs.stringify({user});
    return this.http.get<any>(`api/getMenuItems?${query}`)
  }
  getUser(id, type): any {
    let query = qs.stringify({id, type});
    return this.http.get<any>(`api/getUser?${query}`)
  }
  UploadFile(stud,file): any {
    var fd = new FormData();
    fd.append('file', file);
      fd.append('s', stud);
    fd.append('filetype', 'csv');
    fd.append('f', 'json');
    return this.http.post<any>('api/UploadFile',fd,stud)
  }

}
