import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';

import { Observable, throwError } from 'rxjs';
import { catchError } from 'rxjs/operators';
import { environment } from '../../../../environments/environment';
import { UnidadeModel } from '../models';

@Injectable({
  providedIn: 'root'
})
export class UnidadesService {

  constructor(
    private http: HttpClient
  ) { }

  private readonly BASE_URL = environment.BASE_URL;

  cadastrarUnidade(unidade: UnidadeModel): Observable<any> {
    let params = `/api/unidade/cadastrar`;
    return this.http.post(this.BASE_URL + params, unidade, { withCredentials: true })
      .pipe(
        catchError(
          (error) => {
            return throwError(
              () => error
            )
          }));
  }

  buscarUnidadeID(id: number): Observable<any> {
    let params = `/api/unidade/${id}`;
    return this.http.get(this.BASE_URL + params, { withCredentials: true })
      .pipe(
        catchError(
          (error) => {
            return throwError(
              () => error
            )
          }));
  }

  buscarTodasUnidades(): Observable<any> {
    let params = `/api/unidade`;
    return this.http.get(this.BASE_URL + params, { withCredentials: true })
      .pipe(
        catchError(
          (error) => {
            return throwError(
              () => error
            )
          }));
  }

  deletarUnidade(id: number): Observable<any> {
    let params = `/api/unidade/${id}`;
    return this.http.delete(this.BASE_URL + params, { withCredentials: true })
      .pipe(
        catchError(
          (error) => {
            return throwError(
              () => error
            )
          }));
  }

  atualizarUnidade(id: number, unidade: UnidadeModel): Observable<any> {
    let params = `/api/unidade/${id}`;
    return this.http.put(this.BASE_URL + params, unidade, { withCredentials: true })
      .pipe(
        catchError(
          (error) => {
            return throwError(
              () => error
            )
          }));
  }
}
