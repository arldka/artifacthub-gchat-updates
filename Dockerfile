FROM gcr.io/distroless/static-debian11:nonroot
WORKDIR /
COPY artifacthub-gchat-updates /
USER nonroot:nonroot
EXPOSE 8080
ENTRYPOINT [ "/artifacthub-gchat-updates" ]