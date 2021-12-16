import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';

import { Observable, throwError } from 'rxjs';
import { catchError } from 'rxjs/operators';
import { environment } from '../../environments/environment';

@Injectable({
  providedIn: 'root'
})
export class UsuarioService {

  constructor(
    private http: HttpClient
  ) { }

  private readonly BASE_URL = environment.BASE_URL;

  buscarUsuario(): Observable<any> {
    let params = `/api/usuario/meus-dados`;
    return this.http.get(this.BASE_URL + params, { withCredentials: true })
      .pipe(
        catchError(
          (error) => {
            return throwError(
              () => error
            )
          }));
  }

  buscarTodosUsuarios(): Observable<any> {
    let params = `/api/usuario`;
    return this.http.get(this.BASE_URL + params, { withCredentials: true })
      .pipe(
        catchError(
          (error) => {
            return throwError(
              () => error
            )
          }));
  }

  uploadFotoPerfil(id: number, arquivo: any): Observable<any> {
    let params = `/api/usuario/foto-perfil/${id}`;
    return this.http.put(this.BASE_URL + params, arquivo, { withCredentials: true })
      .pipe(
        catchError(
          (error) => {
            return throwError(
              () => error
            )
          }));
  }

  downloadFotoPerfil(id: number): Observable<any> {
    let params = `/api/usuario/foto-perfil/${id}`;
    return this.http.get(this.BASE_URL + params, { withCredentials: true })
      .pipe(
        catchError(
          (error) => {
            return throwError(
              () => error
            )
          }));
  }

}
