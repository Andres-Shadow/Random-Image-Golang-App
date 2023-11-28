FROM debian
WORKDIR /app
COPY srvimgdocker .
COPY fotos ./fotos
EXPOSE 7900
ENTRYPOINT ["./srvimgdocker"]
CMD ["7900", "fotos"]