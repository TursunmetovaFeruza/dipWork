import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';
import { HomeComponent } from './home/home.component';
import { BarComponent } from './bar/bar.component';
import { LoginComponent } from './login/login.component'
import { StudentMainComponent } from './student-main/student-main.component';
import { UsersComponent } from './users/users.component';
import { MainMenuComponent } from './main-menu/main-menu.component';
import { ProfileComponent } from './profile/profile.component';
import { AttendanceComponent } from './attendance/attendance.component';
import { AddComponent } from './add/add.component';
import { FingerComponent } from './finger/finger.component';

const routes: Routes = [
  {
    path: 'home',
    component: HomeComponent
  },
  {
    path: 'login',
    component: LoginComponent
  },
  {
    path: 'bar',
    component: BarComponent,
    children: [
      {
        path: 'student-main',
        component: StudentMainComponent
      },
      {
        path: 'profile',
        component: ProfileComponent
      },
      {
        path: 'attendance',
        component: AttendanceComponent
      },
      {
        path: ':type',
        component: MainMenuComponent
      },
      {
        path: 'users/:child',
        component: UsersComponent
      },
      {
        path:'add/finger',
        component:FingerComponent
      },
      {
        path: 'add/:child',
        component: AddComponent
      }
    ]
  },
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
