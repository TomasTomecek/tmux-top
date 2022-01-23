%global debug_package   %{nil}
%global provider        github
%global provider_tld    com
%global project         TomasTomecek
%global repo            tmux-top
# https://github.com/TomasTomecek/tmux-top
%global goipath         %{provider}.%{provider_tld}/%{project}/%{repo}
%global forgeurl        https://%{goipath}
%global commit          910ef1f72549a703c3c39abaefefe9a80d0b22fd
%global shortcommit     %(c=%{commit}; echo ${c:0:7})
%global golicenses    LICENSE
%gometa

Name:           tmux-top
Version:        0.1.0
Release:        1%{?dist}
Summary:        Monitoring information for your tmux status line.
License:        GPLv2+
URL:            %{gourl}
Source0:        %{gosource}

BuildRequires:  make
BuildRequires:  golang(github.com/urfave/cli)


%description
Monitoring information for your tmux status line.

tmux-top allows you to see your:

 * load
 * memory usage
 * network information
 * I/O


%prep
%autosetup -n %{name}-%{commit}

%build
%gobuild -o %{gobuilddir}/bin/%{name} %{goipath}/cmd/tmux-top


%install
install -m 0755 -vd %{buildroot}%{_bindir}
install -m 0755 -vp %{gobuilddir}/bin/tmux-top %{buildroot}%{_bindir}/


%check
export GOPATH=$(pwd):%{gopath}
make test


%files
%doc README.md
%license LICENSE
%{_bindir}/%{name}


%changelog
* Sun Feb 03 2019 Fedora Release Engineering <releng@fedoraproject.org> - 0.0.4-4
- Rebuilt for https://fedoraproject.org/wiki/Fedora_30_Mass_Rebuild

* Sat Jul 14 2018 Fedora Release Engineering <releng@fedoraproject.org> - 0.0.4-3
- Rebuilt for https://fedoraproject.org/wiki/Fedora_29_Mass_Rebuild

* Fri Feb 09 2018 Fedora Release Engineering <releng@fedoraproject.org> - 0.0.4-2
- Rebuilt for https://fedoraproject.org/wiki/Fedora_28_Mass_Rebuild

* Sun Oct 29 2017 Tomas Tomecek <ttomecek@redhat.com> - 0.0.4-1
- new upstream release 0.0.4

* Sun Oct 29 2017 Tomas Tomecek <ttomecek@redhat.com> - 0.0.3-1
- new upstream release 0.0.3

* Thu Aug 03 2017 Fedora Release Engineering <releng@fedoraproject.org> - 0.0.2-3
- Rebuilt for https://fedoraproject.org/wiki/Fedora_27_Binutils_Mass_Rebuild

* Thu Jul 27 2017 Fedora Release Engineering <releng@fedoraproject.org> - 0.0.2-2
- Rebuilt for https://fedoraproject.org/wiki/Fedora_27_Mass_Rebuild

* Sat Jun 03 2017 Tomas Tomecek <ttomecek@redhat.com> - 0.0.2-1
- new upstream release 0.0.2

* Sat Feb 11 2017 Fedora Release Engineering <releng@fedoraproject.org> - 0.0.1-7
- Rebuilt for https://fedoraproject.org/wiki/Fedora_26_Mass_Rebuild

* Thu Jul 21 2016 Fedora Release Engineering <rel-eng@lists.fedoraproject.org> - 0.0.1-6
- https://fedoraproject.org/wiki/Changes/golang1.7

* Mon Feb 22 2016 Fedora Release Engineering <rel-eng@lists.fedoraproject.org> - 0.0.1-5
- https://fedoraproject.org/wiki/Changes/golang1.6

* Fri Feb 05 2016 Fedora Release Engineering <releng@fedoraproject.org> - 0.0.1-4
- Rebuilt for https://fedoraproject.org/wiki/Fedora_24_Mass_Rebuild

* Fri Jun 19 2015 Fedora Release Engineering <rel-eng@lists.fedoraproject.org> - 0.0.1-3
- Rebuilt for https://fedoraproject.org/wiki/Fedora_23_Mass_Rebuild

* Mon Mar 16 2015 Tomas Tomecek <ttomecek@redhat.com> - 0.0.1-2
- add devel subpackage (patch by jchaloup@redhat.com)

* Fri Mar 13 2015 Tomas Tomecek <ttomecek@redhat.com> - 0.0.1-1
- initial release

