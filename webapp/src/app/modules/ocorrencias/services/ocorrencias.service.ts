import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';

import { Observable, throwError } from 'rxjs';
import { catchError } from 'rxjs/operators';
import { environment } from '../../../../environments/environment';
import { OcorrenciaModel } from '../models';

@Injectable({
  providedIn: 'root'
})
export class OcorrenciasService {

  constructor(
    private http: HttpClient
  ) { }

  private readonly BASE_URL = environment.BASE_URL;

  cadastrarOcorrencia(ocorrencia: OcorrenciaModel): Observable<any> {
    let params = `/api/ocorrencia/cadastrar`;
    return this.http.post(this.BASE_URL + params, ocorrencia, { withCredentials: true })
      .pipe(
        catchError(
          (error) => {
            return throwError(
              () => error
            )
          }));
  }

  buscarOcorrenciaID(id: number): Observable<any> {
    let params = `/api/ocorrencia/${id}`;
    return this.http.get(this.BASE_URL + params, { withCredentials: true })
      .pipe(
        catchError(
          (error) => {
            return throwError(
              () => error
            )
          }));
  }

  buscarTodasOcorrencias(): Observable<any> {
    let params = `/api/ocorrencia`;
    return this.http.get(this.BASE_URL + params, { withCredentials: true })
      .pipe(
        catchError(
          (error) => {
            return throwError(
              () => error
            )
          }));
  }

  buscarTodosTipoOcorrencias(): Observable<any> {
    let params = `/api/tipo-ocorrencia`;
    return this.http.get(this.BASE_URL + params, { withCredentials: true })
      .pipe(
        catchError(
          (error) => {
            return throwError(
              () => error
            )
          }));
  }

  buscarStatusDashboard(): Observable<any> {
    let params = `/api/ocorrencia/status`;
    return this.http.get(this.BASE_URL + params, { withCredentials: true })
      .pipe(
        catchError(
          (error) => {
            return throwError(
              () => error
            )
          }));
  }

  deletarOcorrencia(id: number): Observable<any> {
    let params = `/api/ocorrencia/${id}`;
    return this.http.delete(this.BASE_URL + params, { withCredentials: true })
      .pipe(
        catchError(
          (error) => {
            return throwError(
              () => error
            )
          }));
  }

  atualizarOcorrencia(ocorrencia: OcorrenciaModel): Observable<any> {
    let params = `/api/ocorrencia/${ocorrencia.id}`;
    return this.http.put(this.BASE_URL + params, ocorrencia, { withCredentials: true })
      .pipe(
        catchError(
          (error) => {
            return throwError(
              () => error
            )
          }));
  }
}