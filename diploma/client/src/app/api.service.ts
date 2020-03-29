import { Injectable } from '@angular/core';
import { HttpClient, HttpHeaders } from '@angular/common/http';

import * as _ from "lodash";



@Injectable()
export class ApiService {

  constructor(private http: HttpClient,) {
  }
  getUsers() {
    return this.http.get<any>('api/users');
  }
  createUsers(value):any {
    return this.http.post<any>('api/users',value);
  }
  signIn(user):any{
    return this.http.post<any>('api/user',user)
  }
 
}
