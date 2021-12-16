import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';

import { Observable, throwError } from 'rxjs';
import { catchError } from 'rxjs/operators';
import { environment } from '../../../../environments/environment';
import { BlocoModel } from '../models';

@Injectable({
  providedIn: 'root'
})
export class BlocosService {

  constructor(
    private http: HttpClient
  ) { }

  private readonly BASE_URL = environment.BASE_URL;

  cadastrarBloco(bloco: BlocoModel): Observable<any> {
    let params = `/api/bloco/cadastrar`;
    return this.http.post(this.BASE_URL + params, bloco, { withCredentials: true })
      .pipe(
        catchError(
          (error) => {
            return throwError(
              () => error
            )
          }));
  }

  buscarBlocoID(id: number): Observable<any> {
    let params = `/api/bloco/${id}`;
    return this.http.get(this.BASE_URL + params, { withCredentials: true })
      .pipe(
        catchError(
          (error) => {
            return throwError(
              () => error
            )
          }));
  }

  buscarTodosBlocos(): Observable<any> {
    let params = `/api/bloco`;
    return this.http.get(this.BASE_URL + params, { withCredentials: true })
      .pipe(
        catchError(
          (error) => {
            return throwError(
              () => error
            )
          }));
  }

  deletarBloco(id: number): Observable<any> {
    let params = `/api/bloco/${id}`;
    return this.http.delete(this.BASE_URL + params, { withCredentials: true })
      .pipe(
        catchError(
          (error) => {
            return throwError(
              () => error
            )
          }));
  }

  atualizarBloco(id: number, bloco: BlocoModel): Observable<any> {
    let params = `/api/bloco/${id}`;
    return this.http.put(this.BASE_URL + params, bloco, { withCredentials: true })
      .pipe(
        catchError(
          (error) => {
            return throwError(
              () => error
            )
          }));
  }
}
