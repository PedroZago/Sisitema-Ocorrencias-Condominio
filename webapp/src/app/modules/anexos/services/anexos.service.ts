import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';

import { Observable, throwError } from 'rxjs';
import { catchError } from 'rxjs/operators';
import { environment } from '../../../../environments/environment';
import { TipoAnexoModel, AnexoModel } from '../models';

@Injectable({
  providedIn: 'root'
})
export class AnexosService {

  constructor(
    private http: HttpClient
  ) { }

  private readonly BASE_URL = environment.BASE_URL;

  cadastrarAnexo(id: number, arquivo: any): Observable<any> {
    let params = `/api/anexo/cadastrar/${id}`;
    return this.http.post(this.BASE_URL + params, arquivo, { withCredentials: true })
      .pipe(
        catchError(
          (error) => {
            return throwError(
              () => error
            )
          }));
  }

  buscarTodosTipoAnexos(): Observable<any> {
    let params = `/api/tipo-anexo`;
    return this.http.get(this.BASE_URL + params, { withCredentials: true })
      .pipe(
        catchError(
          (error) => {
            return throwError(
              () => error
            )
          }));
  }

  buscarTodosAnexosPorOcorrencia(id: number): Observable<any> {
    let params = `/api/anexo/ocorrencia/${id}`;
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
